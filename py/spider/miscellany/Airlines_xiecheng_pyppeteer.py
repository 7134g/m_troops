import asyncio
from pyppeteer import launch

async def main(tasks):
    try:
        for task in tasks:
            browser = await creat()
            page = await browser.newPage()


            task["msg"] = "0000"
            airline = task['company']
            price = 0
            # https://flights.ctrip.com/international/search/oneway-{start}-{end}?depdate={date}directflight=1
            # https://flights.ctrip.com/international/search/oneway-{start}-{end}?depdate={date}&cabin=y_s&adult=1&child=0&infant=0&directflight=1
            url = "https://flights.ctrip.com/international/search/oneway-{start}-{end}?depdate={date}&cabin=y_s&adult=1&child=0&infant=0&directflight=1&airline={airline}".format(
                start=task['start'],
                end=task['end'],
                date=task['date'],
                airline=airline[:2])

            await page.setViewport({'width': 1366, 'height': 768})
            await page.goto(url)
            while not await page.xpath("//div[contains(@id, 'comfort-{}')]".format(airline)):
                pass

            tag_index = -1  # 用于判断是否有该航线
            place_list = await page.xpath('//span[@class="plane-No"]')
            for index, place in enumerate(place_list):
                place = await (await (place.getProperty("textContent"))).jsonValue()
                print(place)
                if task['company'] in str(place):
                    tag_index = index
                    break


            if tag_index != -1:
                # 价格
                while not await page.xpath('//div[@class="price-box"]/div'):
                    pass
                price_list = await page.xpath('//div[@class="price-box"]/div')
                price = (await (await price_list[tag_index].getProperty("textContent")).jsonValue())[1:]
                print("当前价格{},价格区间{}-{}".format(price, task['min_price'], task['max_price']))
                # 是否没超出价格
                if not (int(task['min_price']) < int(price) < int(task['max_price'])):
                    break

            if price:
                pay_id = "#{}_0".format(tag_index)
                await page.click(pay_id)
                await asyncio.gather(
                    page.waitForNavigation(),
                    page.click(pay_id, clickOptions),
                )



    finally:
        await browser.close()
        print("close")

async def creat():
    launch_kwargs = {
        # 控制是否为无头模式
        "headless": False,
        # chrome启动命令行参数
        "args": [
            '--window-size=1366,850',
            # 不显示信息栏  比如 chrome正在受到自动测试软件的控制 ...
            "--disable-infobars",
            # log等级设置 在某些不是那么完整的系统里 如果使用默认的日志等级 可能会出现一大堆的warning信息
            "--log-level=3",
            # 设置UA
            "--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36",
        ],
        # 用户数据保存目录 这个最好也自己指定一个目录
        # 如果不指定的话，chrome会自动新建一个临时目录使用，在浏览器退出的时候会自动删除临时目录
        # 在删除的时候可能会删除失败（不知道为什么会出现权限问题，我用的windows） 导致浏览器退出失败
        # 然后chrome进程就会一直没有退出 CPU就会狂飙到99%
        "userDataDir": r"D:\FXJprograme\FXJ\test\temp",
    }
    browser = await launch({'headless': False})
    return browser

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
asyncio.get_event_loop().run_until_complete(main(params))