# -*- coding: utf-8 -*-
# Scrapy settings for jingdong project
# For simplicity, this file contains only settings considered important or
# commonly used. You can find more settings consulting the documentation:
#
#     http://doc.scrapy.org/en/latest/topics/settings.html
#     http://scrapy.readthedocs.org/en/latest/topics/downloader-middleware.html
#     http://scrapy.readthedocs.org/en/latest/topics/spider-middleware.html
from fake_useragent import UserAgent
ua = UserAgent()


BOT_NAME = 'jingdong'

SPIDER_MODULES = ['jingdong.spiders']
NEWSPIDER_MODULE = 'jingdong.spiders'


MONGO_URL = 'localhost'
MONGO_DB = 'jingdong'

KEYWORDS = "图书"

# Crawl responsibly by identifying yourself (and your website) on the user-agent


USER_AGENT = ua.random

# Obey robots.txt rules

ROBOTSTXT_OBEY = False

# 请填写2~98页
PAGE_NUM = 2



# LOG_LEVEL= 'WARNING'

# Configure maximum concurrent requests performed by Scrapy (default: 16)
#CONCURRENT_REQUESTS = 32

# Configure a delay for requests for the same website (default: 0)
# See http://scrapy.readthedocs.org/en/latest/topics/settings.html#download-delay
# See also autothrottle settings and docs
#DOWNLOAD_DELAY = 3
# The download delay setting will honor only one of:
#CONCURRENT_REQUESTS_PER_DOMAIN = 16
#CONCURRENT_REQUESTS_PER_IP = 16

# Disable cookies (enabled by default)
#COOKIES_ENABLED = False

# Disable Telnet Console (enabled by default)
#TELNETCONSOLE_ENABLED = False

# Override the default request headers:
DEFAULT_REQUEST_HEADERS = {
  'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8',
  'Accept-Language': 'zh-CN,zh;q=0.9',
    'user-agent': USER_AGENT,
    'Connection': 'keep - alive',
    'referer': 'https://search.jd.com/Search?keyword=%E5%9B%BE%E4%B9%A6&enc=utf-8&wq=%E5%9B%BE%E4%B9%A6&page=1'

}

# Enable or disable spider middlewares
# See http://scrapy.readthedocs.org/en/latest/topics/spider-middleware.html
#SPIDER_MIDDLEWARES = {
#    'jingdong.middlewares.JingdongSpiderMiddleware': 543,
#}

# Enable or disable downloader middlewares
# See http://scrapy.readthedocs.org/en/latest/topics/downloader-middleware.html
#DOWNLOADER_MIDDLEWARES = {
#    'jingdong.middlewares.MyCustomDownloaderMiddleware': 543,
#}

# Enable or disable extensions
# See http://scrapy.readthedocs.org/en/latest/topics/extensions.html
#EXTENSIONS = {
#    'scrapy.extensions.telnet.TelnetConsole': None,
#}

# Configure item pipelines
# See http://scrapy.readthedocs.org/en/latest/topics/item-pipeline.html
ITEM_PIPELINES = {
    # 'jingdong.pipelines.JingdongPipeline': 300,
    # 'jingdong.pipelines.MongodbPipeline': 400,
    'scrapy_redis.pipelines.RedisPipeline': 301,
}

# 分布式组件
# 修改调度器
SCHEDULER = "scrapy_redis.scheduler.Scheduler"
#
# 去除重复
DUPEFILTER_CLASS = "scrapy_redis.dupefilter.RFPDupeFilter"

# redis的连接信息
# REDIS_URL = 'redis://用户名:密码@服务器公网IP:端口号'
REDIS_URL = 'redis://Administrator:test12@39.108.167.148:6379'

# 可以暂停爬取，保留已爬取的全部url及指纹
SCHEDULER_PEaRSIST = True

# # 启动时是否清空爬取队列
# SCHEDULER_FLUSH_ON_START = True


# Enable and configure the AutoThrottle extension (disabled by default)
# See http://doc.scrapy.org/en/latest/topics/autothrottle.html
#AUTOTHROTTLE_ENABLED = True
# The initial download delay
#AUTOTHROTTLE_START_DELAY = 5
# The maximum download delay to be set in case of high latencies
#AUTOTHROTTLE_MAX_DELAY = 60
# The average number of requests Scrapy should be sending in parallel to
# each remote server
#AUTOTHROTTLE_TARGET_CONCURRENCY = 1.0
# Enable showing throttling stats for every response received:
#AUTOTHROTTLE_DEBUG = False

# Enable and configure HTTP caching (disabled by default)
# See http://scrapy.readthedocs.org/en/latest/topics/downloader-middleware.html#httpcache-middleware-settings
#HTTPCACHE_ENABLED = True
#HTTPCACHE_EXPIRATION_SECS = 0
#HTTPCACHE_DIR = 'httpcache'
#HTTPCACHE_IGNORE_HTTP_CODES = []
#HTTPCACHE_STORAGE = 'scrapy.extensions.httpcache.FilesystemCacheStorage'



# class transCookie:
#     def __init__(self, cookie):
#         self.cookie = cookie
#
#     def stringToDict(self):
#         '''
#         将从浏览器上Copy来的cookie字符串转化为Scrapy能使用的Dict
#         :return:
#         '''
#         itemDict = {}
#         items = self.cookie.split(';')
#         for item in items:
#             key = item.split('=')[0].replace(' ', '')
#             value = item.split('=')[1]
#             itemDict[key] = value
#         return itemDict
#
# COOKIE = transCookie('')
# COOKIE_CLECK = COOKIE.stringToDict()