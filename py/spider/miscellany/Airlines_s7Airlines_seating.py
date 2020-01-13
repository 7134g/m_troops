import asyncio
import aiohttp
import traceback
import time
from pprint import pprint
from lxml import etree
from aiohttp import ClientError

from _decorator import timer


@timer
async def Do(task: dict, logger) -> list:
    print('gogog')
    start = task["start"]
    end = task["end"]
    company = task["company"]
    date = task["date"]
    userinfo = task['user']

    connector = aiohttp.TCPConnector(verify_ssl=False)
    headers = {
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36'}
    async with aiohttp.ClientSession(connector=connector, headers=headers) as session:
        # TODO 不清楚怎么定位
        url = 'https://ibe.s7-airlines.com/air'
        params = {
            "id": "deeplink",
            "DD1": date,
            "AA1": end,
            "DA1": start,
            "LAN": "zh",
        }
        data_company = {
            "_eventId":"changeSelection",
            # "searchParams.directOnly":"true",
            # "searchParams.redemption":"false",
            # "searchParams.selectedSubsidizedPassengerType":"",
            # "searchParams.tripType":"ONE_WAY",
            # "searchParams.outboundRoute.origin":start,
            # "searchParams.outboundRoute.destination":end,
            # "searchParams.outboundRoute.departureDate":date,
            # "searchParams.passengersAmount.adults":"1",
            # "searchParams.passengersAmount.children":"0",
            # "searchParams.passengersAmount.infants":"0",
            "clientId":"ibe",
            "execution":"e1s1",
            # "ibe_conversation":"",
        }
        data_trip = {
            "_eventId": "proceed",
            "insuranceSupplier": "ALFA",
            "clientId": "ibe",
            "promoCode": '',
            "email": '',
        }
        data_temp = {
            "_eventId": "proceed",
            "options-schedule": 0,
            "promoCode": "",
            "clientId": "ibe"
        }
        order_info = {
            "_eventId": "createOrder",
            "clientId": "ibe",
            "contactsRequest.contacts.emails[0].email": userinfo['email'],
            "contactsRequest.contacts.phones[0].countryCode": 7,
            "contactsRequest.contacts.phones[0].number": userinfo['phone'],
            "passengersRequest.passengers[0].document.dateOfBirth": userinfo['born'],
            "passengersRequest.passengers[0].document.expirationDate": userinfo['passport_time'],
            "passengersRequest.passengers[0].document.nationality": "CN",
            "passengersRequest.passengers[0].document.number": userinfo['passport'],
            "passengersRequest.passengers[0].document.type": "OTHER",
            "passengersRequest.passengers[0].gender": "MALE",
            "passengersRequest.passengers[0].id": "",
            "passengersRequest.passengers[0].index": "1",
            "passengersRequest.passengers[0].lead": "TRUE",
            "passengersRequest.passengers[0].loyaltyProgram.programOwner": "S7",
            "passengersRequest.passengers[0].name.firstName": userinfo['surnames'],
            "passengersRequest.passengers[0].name.lastName": userinfo['name'],
            "passengersRequest.passengers[0].profile": "FALSE",
            "passengersRequest.passengers[0].type": "ADT",
            "paymentRequest.agreedWithVisaBlock": "TRUE",
            "paymentRequest.agreePersonalDataProcessing": "TRUE",
            "paymentRequest.paymentMethodType": "LATER",
            "radio-payment": "on",
        }
        try:
            async with session.get(url=url, params=params) as response:
                base_url = str(response.url)
                print(base_url)
                ibe_conversation = base_url.split("=")[-1]
                # print(ibe_conversation)
                html = etree.HTML(await response.text())
                order_tesk = html.xpath('//div/@data-option-set-id')[0]
                # print(order_tesk)
                company_list = html.xpath(
                    '//div[@class="additional-info"]//span[@class="number js-flight-number"]/@data-id')
                company_index = company_list.index(company)
                # print(company_list)
                nodes = html.xpath('//div[@class="select-item-simple"]')
                price_list = nodes[company_index].xpath('.//span[@class="js-currency-amount"]/@data-amount')
                # print(price_list)
                price_index = -1
                for i,v in enumerate(price_list):
                    if (int(task['min_price']) < int(v) < int(task['max_price'])):
                        price_index = i
                        price = v
                        break
                if price_index == -1:
                    raise asyncio.CancelledError("不在价格范围")

                company_task = nodes[company_index].xpath('.//input[@data-qa="select-option"]/@data-option-id')
                # print(company_task[company_index])
                price_task = nodes[company_index].xpath('.//input[@data-qa="select-option"]/@data-solution-id')
                # print(price_task[price_index])

                data_company["selection.optionSetId"] = order_tesk
                data_company["selection.optionId"] = company_task[company_index]
                data_company["selection.solutionId"] = price_task[price_index]
                data_company["ibe_conversation"] = ibe_conversation


            async with session.post("https://ibe.s7-airlines.com/air",data=data_company) as response_company:
                # print(await response_company.text())
                pass


            async with session.post(url=base_url, data=data_trip) as response_trip:
                html = etree.HTML(await response_trip.text())
                # print(await response_trip.text())
                passengers_id = html.xpath('//div[@class="item js_item"]/@data-passenger-id')[0]
                price_trip = html.xpath('//span[@class="price-amount"]/text()')[0]
                print(price_trip)
                order_info["passengersRequest.passengers[0].id"] = passengers_id

            async with session.post(url=base_url, data=data_temp) as response_temp:
                # pprint((await response_temp.text())[:150])
                pass

            async with session.post(url=base_url, data=order_info) as response_seat:
                '''获取cookies'''
                # pprint((await response_seat.text())[:150])
                cookies = dict(session.cookie_jar.filter_cookies(base_url))
                key_result = "".join([date.replace('-', ''), start, end, company])
                result = [
                    key_result, {
                        "order": ibe_conversation,
                        "cookies": cookies,
                        "modele_name": "1.2.1",
                        "type": 2,
                        "parent": task['ts'],
                        "up": int(time.time() * 1000),
                        "baseTime": 86400000,
                    },
                    int(price.replace(',',''))
                ]
                return result

        except Exception:
            # TODO 非异步，会阻塞
            logger.error("占座操作失败, 错误信息如下{tp}".format(tp=traceback.format_exc()))
            print("占座操作失败, 错误信息如下{tp}".format(tp=traceback.format_exc()))
            pass

# 逐个取出
# async def run_for(params: list)-> list:
#     tasks = [asyncio.ensure_future(post_order(param)) for param in params]
#     for task in asyncio.as_completed(tasks):
#         result = await task
#         print(result)



async def main_ctrl(params: list, logger) -> list:
    print("开始s7占座")
    results = []
    failure_task = []
    for task in params:
        results.append(await Do(task, logger))

    return [results, failure_task]
    # loop = asyncio.get_event_loop()
    # tasks = [asyncio.ensure_future(post_order(task, logger)) for task in params]
    # loop.run_until_complete(asyncio.wait(tasks))
    # results = [task.result() for task in tasks]
    # failure_task = []
    # return [results, failure_task]
    # print("执行完毕")


if __name__ == '__main__':
    import logmanage
    test_logger = logmanage.get_log("我是测试占座日志")
    test_params = [
        {
            'up': 1234567890,
            'type': 1,
            "modele_name": '1.1.1',
            'start': 'PEK',
            'end': 'OVB',
            'company': 'S75722',
            'date': '2019-08-31',
            'min_price': 0,
            'max_price': 40000,
            "ts": 123123,
            'user': {
                'surnames': 'LANA',
                'name': 'WA',
                'gender': 'M',
                'country': '中国大陆',
                'passport': 'EC565987',
                'born': '30.12.1986',
                'passport_time': '11.11.2030',
                'phone': '13944688866',
                'email': '13944688866@163.com',
            }
        },

    ]


    loop = asyncio.get_event_loop()
    tasks = [asyncio.ensure_future(Do(task, test_logger)) for task in test_params]
    loop.run_until_complete(asyncio.wait(tasks))
    results = [task.result() for task in tasks]
    failure_task = []
    print([results, failure_task])
