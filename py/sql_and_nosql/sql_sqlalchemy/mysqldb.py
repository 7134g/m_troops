import datetime
from dateutil.parser import parse
from random import randint
from sqlalchemy import MetaData, Table, select, and_
from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker


MYSQL_URL = ""
# 后面改，父类继承初始化MYSQL_URL

class MysqlClient(object):
    def __init__(self):
        # 寻找Base的所有子类，按照子类的结构在数据库中生成对应的数据表信息
        self.engine = create_engine(MYSQL_URL, encoding='utf-8', pool_size=100)
        # self.connection = self.engine.connect()

    def conn_table(self, table_name):
        """
        第一种连接方式
        :param table_name:
        :return:
        """
        metadata = MetaData(self.engine)
        census = Table(table_name, metadata, autoload=True, autoload_with=self.engine)
        connection = self.engine.connect()
        return census, connection

    def creat_session(self):
        """
        第二种连接方式
        :return:
        """
        # 创建session类型
        DBSession = sessionmaker(bind=self.engine)
        # 创建session对象
        session = DBSession()
        return session

    def data_return(self, data):
        """
        处理mysql对象
        :param data: mysql对象,mysql对象,...
        :return: list[list],str
        """
        if isinstance(data, list):
            for index, value in enumerate(data):
                data[index] = list(value)
            return data
        else:
            return list(data)

    def count(self, table_name):
        """
        数据总数
        :param table_name: 表名
        :return: int
        """
        census, connection = self.conn_table(table_name)
        stmt = select([census])
        count = len(connection.execute(stmt).fetchall())
        connection.close()
        return count

    def index(self, className, session, indexParams):
        obtain = session.query(className).filter(and_(
            *indexParams
        ))
        return obtain

    def updata(self, className, indexParams, pack_dict=None):
        session = self.creat_session()
        obtain = self.index(className, session, indexParams)
        if pack_dict:
            result = obtain.update(pack_dict)
            session.commit()
        else:
            result = 0

        session.close()
        return result

    def r_timestamp(self, table_name, timer=0):
        """
        读取mysql_seatpost任务
        :param table_name: 表名
        :param timer: 时间
        :return: list[list],str
        """
        task_result = []
        census, connection = self.conn_table(table_name)
        stmt = select([census])
        stmt = stmt.where(census.columns.ts > timer)
        data_list = connection.execute(stmt).fetchall()
        # 筛选超过今天的出发日期
        for data in data_list:
            start_time = parse(data[3])
            now_time = datetime.datetime.now()
            if now_time < start_time:
                task_result.append(data)

        connection.close()
        return self.data_return(task_result)

    def r_all(self, table_name):
        """
        返回全部数据
        table_name 表名字
        :return: list[list],str
        """
        census, connection = self.conn_table(table_name)
        stmt = select([census])
        results = connection.execute(stmt).fetchall()
        connection.close()
        return self.data_return(results)

    def r_area(self, table_name, area: int):
        """
        随机获取容量范围内的部分
        :param table_name: 表名
        :param area: 获取资源数
        :return: list[list],str
        """
        census, connection = self.conn_table(table_name)
        stmt = select([census])
        lenght = self.count(table_name)

        helf_lenght = int(lenght / 2)
        if area < helf_lenght:
            index = randint(area+1, helf_lenght)
        elif area == lenght:
            return self.r_all(table_name)
        else:
            index = randint(area+1, lenght)

        stmt = stmt.where(and_(census.columns.id >= index-area, census.columns.id < index))
        results = connection.execute(stmt).fetchall()
        # print(" 起{} -终{} -得{} -求{}".format(index-area, index, str(len(results)), area))
        connection.close()
        return self.data_return(results)

    def r_absolute_area(self, table_name, area: int):
        """
        动态申请需要msg参数，不足时自我复制
        :param table_name: 表名
        :param area: 获取数
        :return: list[list],str
        """
        lenght = self.count(table_name)  # 获取数据库总个数
        residue = area % lenght
        # census = self.conn_table(table_name)
        # stmt = select([census])

        if area <= lenght:
            return  self.r_area(table_name, area)

        elif area > lenght and area < lenght*2:
            first_part = self.r_all(table_name)
            second_part = first_part[:residue]
            return first_part + second_part

        else:
            copy_count = area // lenght
            data1 = []
            all_area = self.r_all(table_name)
            for _count in range(copy_count):
                data1 += all_area

            data2 = all_area[:residue]
            return data1+data2

    def r_choice_list(self, table_name, count: int):
        """
        随机选择一批数据
        :param table_name:
        :param count:
        :return: list[list],str
        """
        lenght = self.count(table_name)

        num_list = []
        census, connection = self.conn_table(table_name)
        for x in range(count):
            stmt = select([census])
            num = randint(0, lenght)
            stmt = stmt.where(census.columns.id == num)
            result = connection.execute(stmt).fetchall()
            num_list += self.data_return(result)

        connection.close()
        return num_list

    def r_index_ts(self, table_name, num: int):
        """
        根据索引选一条
        :param table_name: 表名
        :return:
        """
        # lenght = self.count(table_name)
        census, connection = self.conn_table(table_name)
        stmt = select([census])
        stmt = stmt.where(census.columns.ts == num)
        result = connection.execute(stmt).fetchall()
        connection.close()
        return self.data_return(result)

    def w_data(self, data):
        """
        SeatPost(start=start,...)
        :param data:
        :return:
        """
        session = self.creat_session()
        session.add(data)
        session.commit()
        session.close()

    def d_once(self, table_name, base_str, num: int):
        connection = self.engine.connect()
        meta = MetaData(bind=self.engine)
        tb_user = Table(table_name, meta, autoload=True, autoload_with=self.engine)
        dlt = tb_user.delete().where(tb_user.columns[base_str] == num)
        # 执行delete操作
        result = connection.execute(dlt)
        # 显示删除的条数
        connection.close()
        return result.rowcount

    def d_SeatCount(self, className, indexParams):
        session = self.creat_session()
        obtain = self.index(className, session, indexParams).delete()
        if obtain:
            session.commit()
            session.close()
        else:
            print("没有seatCount为0的数据")

    def r_seleter_once(self, className, indexParams):
        """
        根据索引选一条
        :param className: 模型名
        :param indexParams: 筛选，indexParams = [ViewData.start == start,....]
        :return: 值
        """
        # lenght = self.count(table_name)
        session = self.creat_session()
        obtain = self.index(className, session, indexParams).all()
        session.close()
        return obtain


class SeatPost:
    def __init__(self,start, end, company, date):
        self.start = start
        self.end = end
        self.company = company
        self.date = date
    pass

if __name__ == '__main__':
    # [103, 'ZHANG/YIDI ', '男', '中国 ', 'E57013072', datetime.datetime(1992, 10, 2, 0, 0), datetime.datetime(2025, 8, 10, 0, 0), '中国 ']
    import time

    timer = int(time.time()*1000)
    num = 2
    m = MysqlClient()
    user = SeatPost(start="ctu",end="tyo",company="CZ4081",date="2019-07-19")
    pack = {"seatcount":200}
    result = m.r_seleter_once(SeatPost,)
    print(result)




