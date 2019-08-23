import redis
import pymongo
import json

def process_item():
    # db=数据库位置，port=端口号， host=ip地址
    rediscli = redis.Redis(host="127.0.0.1",port=6379, db =0)
    
    mongocli = pymongo.MongoClient(host"127.0.0.1",post=27017)


    dbname = mongocli['数据库名称']

    sheetname = dbname["数据库表名"]

    offset = 0

    while True:
        # redis 数据表名和数据 ,将数据从redis里pop出来
        source, data = rediscli.blpop('goods:items')
        offset += 1
        # 将json格式转化为字典
        data = json.loads(data)
        dbname[sheetname].update({'title':data['title']},{'$set':data},True)
        print(offset)

    mongocli.close()

    
if __name__ == '__main__':
    process_item()
