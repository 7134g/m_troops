from copy import copy

from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver import ChromeOptions, Chrome
from pprint import pprint
import traceback
import json
import time

from _decorator import timer

MISS_ERROR = 3


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


def check_driver(driver):
    try_max = 0
    while "未连接到互联网" in driver.page_source:
        driver.quit()
        driver = SeleniumDriver().Chrome()
        driver.maximize_window()
        try_max += 1
        if try_max > 10:
            raise Exception("ip全部过期啦~！")

    return driver


@timer
def Do(driver, param_list, error, logger):
    result = "0000"
    for param in param_list:
        try:
            login_url = "https://flights.ctrip.com"
            driver.get(login_url)
            driver = check_driver(driver)
            # 添加cookies信息
            driver.delete_all_cookies()  # 清空原来的
            cookies = param['cookies']
            # cookies = json.loads(param['cookies'])
            for cookie in cookies:
                try:
                    # cookie["expiry"] = int(cookie.get("expiry",0))
                    driver.add_cookie(cookie)
                except:
                    traceback.print_exc()
                    pass

            url = "https://flights.ctrip.com/online/orderdetail/index?oid={}".format(param['order'])
            driver.get(url)
            driver = check_driver(driver)
            element = WebDriverWait(driver, 3, 0.5)

            # 定位取消按键
            element.until(EC.presence_of_element_located((By.XPATH, '//a[@data-ubt-t="btns-1005"]'))).click()

            # 确认取消
            element.until(EC.presence_of_element_located((By.CLASS_NAME, 'ant-modal-footer')))
            driver.find_element(By.XPATH, '/html/body/div[9]/div/div[2]/div/div[2]/div[2]/div/div[2]/div').click()

            if "取消失败" in driver.page_source:
                raise Exception("取消失败")



            result = {
                "redisKeyName":param["redisKeyName"],
                "parent":param["parent"],
                "up":int(time.time()*1000),
            }
            # pprint("取消成功")
            yield result
            driver.delete_all_cookies()

        except:
            driver.delete_all_cookies()
            error.append(param)
            traceback.print_exc()
            logger.error("取消占位失败{},{}".format(json.dumps(param), traceback.format_exc()))
            yield result

# def main_ctrl(params, logger):
#     if params['type'] != 2:
#         raise Exception('任务错误')
#     if params["modele_name"] != 'plat.xiecheng.xiecheng_cencel':
#         raise  Exception('模块错误')
#
#     return Do(params, logger)

# def main_ctrl(driver, params_list, logger):
def main_ctrl(params_list, logger):
    results = []
    error = []
    failure_task = []
    try:
        driver = SeleniumDriver().Chrome()
        driver.maximize_window()
        for count in range(MISS_ERROR):
            try:
                if count == 0:
                    for result in Do(driver,params_list, error, logger):
                        if isinstance(result,dict):
                            logger.info('{module_name}.py 执行成功:{jsondata}'.format(
                                module_name=params_list[0]["modele_name"],
                                jsondata=json.dumps(result, ensure_ascii=False)))
                            results.append(result)
                        else:
                            if result=="0000":
                                raise Exception
                            # str,删除无效的占座任务 TODO
                            if result:
                                failure_index = []
                                for index, failure in enumerate(params_list):
                                    if failure["redisKeyName"] in result:
                                        failure_index.append(index)
                                for index in failure_index[::-1]:
                                    params_list.pop(index)
                                failure_task.append(result)

                elif error:
                    repeat_task = copy(error)
                    error = []
                    for result in Do(driver, repeat_task, error, logger):
                        # 记录日志
                        if isinstance(result, dict):
                            logger.info('{module_name}.py 执行成功:{jsondata}'.format(
                                module_name=params_list[0]["modele_name"],
                                jsondata=json.dumps(result, ensure_ascii=False)))
                            results.append(result)
                        if result == "0000":
                            raise Exception


                print("执行完毕，开始处理错误任务")
                print("当前第{}次执行".format(count))
                if len(results)>=len(params_list):
                    print("已达到任务要求数量，停止处理错误任务回发操作")

            except:
                traceback.print_exc()
                driver.quit()
                driver = SeleniumDriver().Chrome()
                driver.maximize_window()


    except:
        logger.error('启动模拟器失败，错误信息{tp}'.format(tp=traceback.format_exc()))

    finally:
        driver.quit()

    return results,failure_task




if __name__ == '__main__':
    import os, psutil
    import logmanage

    # driver = SeleniumDriver().Chrome()
    # driver.maximize_window()
    process = psutil.Process(os.getpid())
    print('Used Memory:', process.memory_info().rss / 1024 / 1024, 'MB')
    logger = logmanage.get_log('测试携程取消座位操作')

    params_list = [{
        'order':"10190701804",
        "cookies":"""[{'domain': '.ctrip.com', 'expiry': 1564548525, 'httpOnly': False, 'name': '_bfs', 'path': '/', 'secure': False, 'value': '1.6'}, {'domain': '.ctrip.com', 'expiry': 1627618725, 'httpOnly': False, 'name': '_bfa', 'path': '/', 'secure': False, 'value': '1.1564546696689.2ciive.1.1564546696689.1564546696689.1.6'}, {'domain': '.ctrip.com', 'expiry': 1577807999.432394, 'httpOnly': True, 'name': '_fpacid', 'path': '/', 'secure': False, 'value': '09031149110241798976'}, {'domain': '.ctrip.com', 'expiry': 1564557508.001677, 'httpOnly': False, 'name': 'IsNonUser', 'path': '/', 'secure': False, 'value': 'T'}, {'domain': '.ctrip.com', 'expiry': 1564557508.001605, 'httpOnly': False, 'name': 'DUID', 'path': '/', 'secure': False, 'value': 'u=95954DFB6C72FD226FAFF6EBE8AF37A2&v=0'}, {'domain': '.ctrip.com', 'expiry': 1564557508.001517, 'httpOnly': True, 'name': 'ticket_ctrip', 'path': '/', 'secure': False, 'value': 'bJ9RlCHVwlu1ZjyusRi+ypZ7X2r4+yojMGCLdmOoktZjH81N1J89q3JU9tYJa13wq4XH/hWje7gXDrj41qs8uwoOPqctBOncbBbenZ+0UHogizgxB/VDu/tlGvMrcsEBTihegKD+pCvl+TSVN6HxF7PN+r1ptNh1Ml34FvVBL2/ycHOEppzBNXupUG53OHie2ud564YXummCn4WtkKimBgn938UFUwI9U+VHGX4U38R/mjOVz2JRka7JxjYY9fLrOVyjH1ZlQN3wagu3j5NbADAJYDEtisbgyEZj1Ux/wLQ='}, {'domain': '.ctrip.com', 'expiry': 1657858719, 'httpOnly': False, 'name': 'GUID', 'path': '/', 'secure': False, 'value': '09031149110241798976'}, {'domain': '.ctrip.com', 'expiry': 3141346709.001422, 'httpOnly': False, 'name': 'AHeadUserInfo', 'path': '/', 'secure': False, 'value': 'VipGrade=0&VipGradeName=%C6%D5%CD%A8%BB%E1%D4%B1&UserName=&NoReadMessageCount=0'}, {'domain': '.ctrip.com', 'expiry': 1564633123, 'httpOnly': False, 'name': '_gid', 'path': '/', 'secure': False, 'value': 'GA1.2.257013116.1564546700'}, {'domain': '.ctrip.com', 'expiry': 1564557508.001319, 'httpOnly': True, 'name': 'cticket', 'path': '/', 'secure': False, 'value': 'D68D8A6DCBE89C596D330EEC59CA49F0C064DBB5B316A4909B86F7DAED23D41A'}, {'domain': '.ctrip.com', 'expiry': 4042022400, 'httpOnly': False, 'name': '_RSG', 'path': '/', 'secure': False, 'value': 'CjK6_tbj5zE1_ABKaaIWD8'}, {'domain': '.ctrip.com', 'expiry': 1580314699, 'httpOnly': False, 'name': '_jzqco', 'path': '/', 'secure': False, 'value': '%7C%7C%7C%7C%7C1.1619362860.1564546699605.1564546699605.1564546699605.1564546699605.1564546699605.0.0.0.1.1'}, {'domain': '.ctrip.com', 'expiry': 1596082697, 'httpOnly': False, 'name': 'FlightIntl', 'path': '/', 'secure': False, 'value': 'Search=[%22TPE|%E5%8F%B0%E5%8C%97(TPE)|617|TPE|480%22%2C%22OSA|%E5%A4%A7%E9%98%AA(OSA)|219|OSA|540%22%2C%222019-08-09%22]'}, {'domain': '.ctrip.com', 'expiry': 1627618723, 'httpOnly': False, 'name': '_ga', 'path': '/', 'secure': False, 'value': 'GA1.2.1382485369.1564546700'}, {'domain': '.ctrip.com', 'expiry': 1627618699, 'httpOnly': False, 'name': '__zpspc', 'path': '/', 'secure': False, 'value': '9.1.1564546699.1564546699.1%234%7C%7C%7C%7C%7C%23'}, {'domain': '.ctrip.com', 'expiry': 1565151500, 'httpOnly': False, 'name': 'MKT_Pagesource', 'path': '/', 'secure': False, 'value': 'PC'}, {'domain': '.ctrip.com', 'expiry': 1564546759, 'httpOnly': False, 'name': '_gat', 'path': '/', 'secure': False, 'value': '1'}, {'domain': '.ctrip.com', 'expiry': 4042022400, 'httpOnly': False, 'name': '_RGUID', 'path': '/', 'secure': False, 'value': 'ef45b65d-ff19-4b4c-8360-4a49142cafc1'}, {'domain': '.ctrip.com', 'expiry': 4042022400, 'httpOnly': False, 'name': '_RF1', 'path': '/', 'secure': False, 'value': '118.114.164.146'}, {'domain': '.ctrip.com', 'expiry': 4042022400, 'httpOnly': False, 'name': '_RDG', 'path': '/', 'secure': False, 'value': '28800503880bdc2169131caafe0dcb61e6'}, {'domain': '.flights.ctrip.com', 'expiry': 1650946696.200413, 'httpOnly': False, 'name': '_abtest_userid', 'path': '/', 'secure': False, 'value': 'b9dd9266-027c-46c5-ad71-d4def5b47894'}]""",
        'type':2,
        "modele_name":'1.1.2',
        "parent": int(time.time() * 1000),
        "up": int(time.time() * 1000),
    },

    ]
    pprint(main_ctrl(params_list, logger))