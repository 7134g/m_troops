from copy import copy
from pprint import pprint

from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.common.exceptions import UnexpectedAlertPresentException, ElementNotVisibleException, TimeoutException
from selenium.webdriver import ChromeOptions, Chrome

from lxml import etree
import time
import traceback
import json

from _decorator import timer

MISS_ERROR = 3
THEAD_TASK_COUNT = 3


class SeleniumDriver:
    def __init__(self):
        self.options = ChromeOptions()
        self.options.add_argument(
            'user-agent=Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36')
        # self.options.add_argument('--headless')  # 去掉可视窗
        self.options.add_argument('--no-sandbox')
        self.options.add_argument('--disable-gpu')
        self.options.add_argument('--log-level=3')
        # self.options.add_experimental_option('excludeSwitches', ['enable-automation'])
        # self.options.add_experimental_option('debuggerAddress', '127.0.0.1:9222')
        prefs = {"profile.managed_default_content_settings.images": 2}
        self.options.add_experimental_option("prefs", prefs)  # 图片不加载

    # Chrome 浏览器
    def Chrome(self):
        # PROXY_IP = get_proxy()
        # PROXY_IP = proxyclient.get_ip()
        # self.options.add_argument('--proxy-server=http://{}'.format(PROXY_IP))
        driver = Chrome(executable_path="D:/python/Scripts/chromedriver.exe",chrome_options=self.options)
        return driver


def user_form(driver, msg):


    WebDriverWait(driver, 5, 0.1).until(EC.presence_of_element_located((By.XPATH, '//input[@value="M"]')))
    surnames = driver.find_element("xpath", '//div[@data-field="SURNAME"]//input[@class="form-input"]').send_keys(msg['surnames'])
    name = driver.find_element("css selector", "div[data-field='GIVEN_NAME'] input").send_keys(msg['name'])
    born = driver.find_element("css selector", "div[data-field='BIRTHDAY'] input").send_keys(msg['born'])
    gender = driver.find_element("css selector", "div[data-field='GENDER'] input[value='M']").click()

    # 点击国家
    country = driver.find_element("css selector", "div[data-field='NATIONALITY'] input").click()
    WebDriverWait(driver, 5, 0.1).until(EC.element_to_be_clickable((By.CLASS_NAME, 'popover-fuzzylist')))
    country_se = driver.find_element("xpath", "//div[@data-field='NATIONALITY']//ul/li[1]").click()
    # button_js = """$().click()"""
    # driver.execute_script(button_js)


    passport = driver.find_element("css selector", "#cardNum_0_0").send_keys(msg['passport'])
    passport_time = driver.find_element("css selector", "#input_card_limit_time_0_0").send_keys(msg['passport_time'])
    phone = driver.find_element("css selector", "div[data-field='MOBILE'] input").send_keys(msg['phone'])
    phone2 = driver.find_element("css selector", "#contactPhoneNum").send_keys(msg['phone'])
    email = driver.find_element("css selector", "div[data-field='EMAIL'] input").send_keys(msg['email'])

@timer
def Do(driver, task_list, error_list, logger):
    """

    :param task: dict
    :return: list—>[str,dict]
    """
    end_task = 0
    for task in task_list:
        # print(task)
        task["msg"] = "0000"
        userinfo = task['user']
        airline = task['company']
        # https://flights.ctrip.com/international/search/oneway-tpe-osa?depdate=2019-08-09&cabin=y_s&adult=1&child=0&infant=0&directflight=1&airline=D7
        url = "https://flights.ctrip.com/international/search/oneway-{start}-{end}?depdate={date}&cabin=y_s&adult=1&child=0&infant=0&directflight=1&airline={airline}".format(
            start=task['start'],
            end=task['end'],
            date=task['date'],
            airline=airline[:2])
        driver.delete_all_cookies()
        driver.get(url)
        element = WebDriverWait(driver, 30, 0.1)
        driver.implicitly_wait(5)  # 隐式
        # print(driver.page_source)

        try:
            # element.until(EC.presence_of_element_located((By.CSS_SELECTOR,".loading.finish" )))
            startRun = int(time.time())
            while not (etree.HTML(driver.page_source).xpath('//div[@class="loading finish"]/@style')):
                if int(time.time())-startRun>30:
                    error_list.append(task)
                    driver.quit()


            # # 定位航司
            # airline_element = "//div[contains(@id, 'comfort-{}')]".format(airline)
            # element.until(EC.presence_of_element_located((By.XPATH, airline_element)))

            tag_index = -1  # 用于判断是否有该航线
            # 获取当前页面航线名称list
            html = etree.HTML(driver.page_source)
            place_list = html.xpath('//span[@class="plane-No"]/text()')
            print(place_list)
            for index, place in enumerate(place_list):
                if task['company'] in place:
                    tag_index = index
                    break

            if tag_index != -1:
                # 价格
                # element.until(EC.presence_of_element_located((By.XPATH, '//div[@class="price-box"]/div')))
                price_list = html.xpath('//div[contains(@id,"price_{}")]/text()'.format(tag_index))
                # print(price_list)
                # price_list = price[0]
                print("当前价格{},价格区间{}-{}".format(price_list, task['min_price'], task['max_price']))

                if not price_list:
                    raise TimeoutException

                price = None
                pay_index = 0
                for price_index,temp_price in enumerate(price_list):
                    if (int(task['min_price']) < int(temp_price) < int(task['max_price'])):
                        price = temp_price # 得到价格
                        pay_index = price_index # 得到索引
                        break

            else:
                raise ElementNotVisibleException


            if price:
                # 点击预订
                print("".join(["开始预定========",price]))
                WebDriverWait(driver, 3, 0.1).until(EC.element_to_be_clickable((By.ID, '{}_{}'.format(tag_index,pay_index)))).click()
                # print(driver.window_handles)
                if len(driver.window_handles)>1:
                    driver.switch_to_window(driver.window_handles[1])
                    driver.close()
                    driver.switch_to_window(driver.window_handles[0])


                # 打开登录窗
                try:
                    WebDriverWait(driver, 3, 0.1).until(EC.element_to_be_clickable((By.ID, 'nologin'))).click()
                except:
                    # raise Exception("登录窗问题")
                    pass

                try:
                    # 个人信息
                    user_form(driver, userinfo)
                except:
                    # WebDriverWait(driver, 1, 0.1).until(EC.element_to_be_clickable((By.ID, 'outer')))
                    continue


                # 提交订单
                element.until(EC.element_to_be_clickable((By.CLASS_NAME, 'btn-next'))).click()
                # print(driver.page_source)

                startRun = int(time.time())
                while not ("护照:{}".format(task["user"]["passport"]) in driver.page_source ):
                    if int(time.time()) - startRun > 30:
                        driver.quit()
                    findStr = "{}/{}".format(task["user"]["surnames"],task["user"]["name"]).replace(" ","")
                    if ((findStr in driver.page_source) or ("目前该舱位已售完" in driver.page_source)) and (not "护照:{}".format(task["user"]["passport"]) in driver.page_source):
                        yield task
                        end_task = 1
                        break
                # print(task_list)
                if end_task:
                    re_task_list = []
                    for task_ in task_list:
                        # if task_["start"]==task['start'] and task_["end"]==task['end'] and task_["date"]==task['date'] and task_["company"]==task['company']:
                        if task_["ts"] == task['ts']:
                            task_["min_price"] = str(int(price)+1)
                            re_task_list.append(task_)
                        else:
                            re_task_list.append(task_)
                    task_list = re_task_list
                    continue
                        # 获取订单号
                # try:
                #     WebDriverWait(driver, 3, 0.5).until(
                #         EC.element_to_be_clickable((By.CLASS_NAME, 'a.btn.btn-primary'))).click()
                # except:
                #     try:
                #         pass
                #         WebDriverWait(driver, 15, 0.5).until(EC.element_to_be_clickable((By.CSS_SELECTOR, 'div.notice')))
                #     except:
                #         traceback.print_exc()
                #         try:
                #             WebDriverWait(driver, 3, 0.5).until(
                #                 EC.element_to_be_clickable((By.CLASS_NAME, 'a.btn.btn-primary'))).click()
                #         except:
                #             pass
                order = str(driver.current_url).split('/')[-2]
                # print(order)

                # WebDriverWait(driver, 20, 0.5).until(EC.presence_of_element_located((By.CSS_SELECTOR, '.header-wrapper')))
                # 获取cookies
                js = 'window.open("https://flights.ctrip.com/online/orderdetail/index?oid={}");'.format(order)
                driver.execute_script(js)
                driver.switch_to_window(driver.window_handles[1])

                startRun = int(time.time())
                while (("没有查到您的订单信息" in driver.page_source) or ("繁忙" in driver.page_source)):
                    driver.refresh()
                    if int(time.time()) - startRun > 30:
                        error_list.append(task)
                        driver.quit()
                    # try:
                    #     WebDriverWait(driver, 3, 0.1).until(EC.presence_of_element_located((By.XPATH, '//a[@data-ubt-v="pay-去支付"]')))
                    # except:
                    #     pass

                cookies = driver.get_cookies()
                # print(cookies)



                # if not cookies:
                #     driver.refresh()
                #     try:
                #         WebDriverWait(driver, 3, 0.1).until(
                #             EC.presence_of_element_located((By.XPATH, '//a[@data-ubt-t="btns-1005"]')))
                #     except:
                #         pass
                #         cookies = driver.get_cookies()



                key_result = "".join(
                    [task["date"].replace('-', ''), task["start"], task["end"], task["company"]])

                # 返回数据
                result = [
                    key_result, {
                        "order": order,
                        "cookies": cookies,
                        "modele_name": "1.1.2",
                        "type": 2,
                        "parent": task['ts'],
                        "up": int(time.time() * 1000),
                        "baseTime":1200000,
                    },
                    price
                ]
                # print(result)
                driver.close()
                yield result
                driver.switch_to_window(driver.window_handles[0])
                driver.delete_all_cookies()
            else:
                print("不在价格范围========================》")
                task['msg'] = '0001'
                yield task
        except UnexpectedAlertPresentException:
            driver.switch_to.alert.accept()
            driver.delete_all_cookies()
            traceback.print_exc()
            length = len(driver.window_handles)
            if length > 1:
                for index in range(1, length):
                    driver.switch_to_window(driver.window_handles[index])
                    driver.close()
                driver.switch_to_window(driver.window_handles[0])

            error_list.append(task)

        except Exception:
            # 如果有弹出框 点击确定
            driver.delete_all_cookies()
            if alert_is_present(driver):
                driver.switch_to.alert.accept()
            traceback.print_exc()
            length = len(driver.window_handles)
            if length>1:
                for index in range(1,length):
                    driver.switch_to_window(driver.window_handles[index])
                    driver.close()
                driver.switch_to_window(driver.window_handles[0])

            error_list.append(task)


def alert_is_present(driver):
    try:
        alert = driver.switch_to.alert
        alert.text
        return alert
    except:
        return False

def main_ctrl(params, logger):
    # if params['type'] != 1:
    #     raise Exception('任务错误')
    # if params["modele_name"] != 'plat.xiecheng.xiecheng_seating':
    #     raise  Exception('模块错误')
    results = []
    error = []
    failure_task = []
    failure_ts_list =set()
    try:
        driver = SeleniumDriver().Chrome()
        driver.maximize_window()
        for count in range(MISS_ERROR):
            if count == 0:
                for result in Do(driver, params, error, logger):
                    # 记录日志
                    if isinstance(result, list):
                        logger.info('{module_name}.py 执行成功:{jsondata}'.format(
                            module_name=result[1]["modele_name"],
                            jsondata=json.dumps(result, ensure_ascii=False)))

                        results.append(result)
                    else:
                        failure_index = []
                        for index,failure in enumerate(params):
                            if result["ts"] == failure["ts"] and result["ts"] not in failure_ts_list:
                                failure_ts_list.add(result["ts"])
                                failure_index.append(index)
                        if result['msg'] == '0001':
                            if result["ts"] not in failure_ts_list:
                                failure_task.append(result)
                        else:
                            error.append(result)
                        # 弹出这类全部任务
                        # for index in failure_index[::-1]:
                        #     params.pop(index)
                        # failure_task.append(result)

            elif error:
                repeat_task = copy(error)
                error = []
                for result in Do(driver, repeat_task, error, logger):
                    # 记录日志
                    if isinstance(result, list):
                        logger.info('{module_name}.py 执行成功:{jsondata}'.format(
                            module_name=result[1]["modele_name"],
                            jsondata=json.dumps(result, ensure_ascii=False)))
                        results.append(result)
                    else:
                        failure_index = []
                        for index,failure in enumerate(params):
                            if result["ts"] == failure["ts"] and result["ts"] not in failure_ts_list:
                                failure_ts_list.add(result["ts"])
                                failure_index.append(index)
                        if result['msg'] == '0001':
                            if result["ts"] not in failure_ts_list:
                                failure_task.append(result)
                        else:
                            error.append(result)

            # print("执行完毕，开始处理错误任务")
            print("当前第{}次执行".format(count))

            if len(results)>=len(params):
                print("已达到任务包数量，停止回发")


    except:
        traceback.print_exc()
        logger.error('启动模拟器失败，错误信息{tp}'.format(tp=traceback.format_exc()))
    finally:
        driver.quit()

    print("全部任务处理完毕")
    return results,failure_task





if __name__ == '__main__':
    import os, psutil
    import logmanage


    process = psutil.Process(os.getpid())
    print('Used Memory:', process.memory_info().rss / 1024 / 1024, 'MB')
    logger = logmanage.get_log("我是测试携程占座日志")
    params = [
        {
        'up': 1234567890,
        'type': 1,
        "modele_name": '1.1.1',
        'start': 'tpe',
        'end': 'osa',
        'company': 'D7370',
        'date': '2019-08-09',
        'min_price': 0,
        'max_price': 40000,
            "ts": 123123,
        'user': {
            'surnames': 'LAO',
            'name': 'WANG',
            'gender': 'M',
            'country': '中国大陆',
            'passport': 'XS1245378',
            'born': '1996-12-30',
            'passport_time': '2029-11-11',
            'phone': '16644663659',
            'email': '16644663659@163.com',
        }
    },
        {
            'up': 1234567890,
            'type': 1,
            'is_stop': 0,
            "modele_name": '1.1.1',
            'start': 'tpe',
            'end': 'osa',
            'company': 'D7370',
            'date': '2019-08-09',
            'min_price': 0,
            'max_price': 40000,
            "ts": 123123,
            'user': {
                'surnames': 'LAO',
                'name': 'WANG',
                'gender': 'M',
                'country': '中国大陆',
                'passport': 'XS1245378',
                'born': '1996-12-30',
                'passport_time': '2029-11-11',
                'phone': '16644663659',
                'email': '16644663659@163.com',
            }
        },
    ]

    result = main_ctrl(params,logger)
    print(result)
    print("数据长度：{}".format(len(result)))



