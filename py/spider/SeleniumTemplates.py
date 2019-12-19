from selenium.webdriver import ChromeOptions,Chrome,PhantomJS
from selenium.webdriver.common.desired_capabilities import DesiredCapabilities
import time

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

    # selenium2.4 ~ 3.141
    # Chrome 浏览器
    def Chrome(self):
        # PROXY_IP = get_proxy()
        # PROXY_IP = proxyclient.get_ip()
        # self.options.add_argument('--proxy-server=http://{}'.format(PROXY_IP))
        driver = Chrome(executable_path="./chromedriver.exe", chrome_options=self.options)
        return driver

    # selenium3.141
    def FireFox(self):
        from selenium import webdriver
        from selenium.webdriver.firefox.options import Options
        options = Options()
        options.add_argument('-headless')
        options.add_argument('--disable-gpu')  # 禁用GPU加速
        options.set_preference('permissions.default.image', 2)  # 禁止加载图片
        options.add_argument('--window-size=1280,800')  # 设置窗口大小
        browser = webdriver.Firefox(executable_path='./geckodriver.exe',
                                    firefox_options=options)
        return browser

    # PhantomJS 浏览器
    @classmethod
    def PhantomJS(cls):
        dcap = dict(DesiredCapabilities.PHANTOMJS)
        dcap["phantomjs.page.settings.userAgent"] = ("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")
        # proxy = get_proxy()
        SERVICE_ARGS = [
            '--disk-cache=true', # 图片不加载
            '--load-images=false',# 图片不加载
            # '--proxy={}'.format(proxy),  # 设置的代理ip
            # '--proxy-type=http',  # 代理类型
            '--ignore-ssl-errors=true',
        ]
        driver = PhantomJS(executable_path='./geckodriver.exe', desired_capabilities=dcap,
                           service_args=SERVICE_ARGS, service_log_path='./log/ghostdriver.log')
        return driver

    @classmethod
    def wait_load(cls, driver, text, timeout=30):
        now_time = time.time()
        while text not in driver.page_source:
            if time.time() > now_time + timeout:
                print("等待超时")
                return
            time.sleep(1)

        print("获取到: {}".format(text))
        return