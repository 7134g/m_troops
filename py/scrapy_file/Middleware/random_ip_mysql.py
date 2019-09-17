# ./jingdong/jingdong/tools/random_ip_mysql
import MySQLdb
import MySQLdb.cursor
import requests

class GetIP(object):

    def __init__(self):
        self.test_ip = crawler.setting.get('TEST_IP_HTML','')

    @classmethon
    def from_setting(cls,crawler):
        return cls(crawler)

    def delete_ip(self,id):
        delete_sql ='''
                delete from proxy_ip where ip='{}' 
            '''.format(ip)
        cusor.execute(delete_sql)
        conn.commit()
        return True

    def judge_ip(self,ip,post):
        # 判断ip是否可用
        try:
            proxies = {'http':ip+post}
            response = requests.get(self.test_ip,proxies=proxies)
            return True
        except Exception as e:
            print('错误')
            self.delete_ip(ip)
            return False
        else:
            code = response.status_code
            if code>=200 and code <= 300:
                print('有效IP')
                return True
            else:
                print('无效IP')
                return False


    def get_random_ip(self):
        random_sql_ip ='''
                SELECT ip,port FROM proxy_ip
                ORDER BY RAND()
                LIMIT 1
                '''
        result = cursor.execute(random_sql_ip)
        for ip_info in cursor.fetchall():
            ip = ip_info[0]
            port = ip_info[1]
            
            judge_result = self.judge_ip(ip,port)
            if judge_result:
                return 'http://{0}{1}'.format(ip,port)
            else:
                return self.get_random_ip()



# 在setting.py中
TEST_IP_HTML = '爬取的网址'


# 在middlewares.py中
from jingdong.tools.random_ip_mysql import GetIP

class RandomIpMiddleware(object):
    def process_requests(self,request,spider):
        get_ip = GetIP()
        request.meta['proxy'] = get_ip.get_random_ip()