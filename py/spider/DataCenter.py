from datetime import datetime
from dateutil.parser import parse
from enum import Enum
from py.spider.ErrorStr import *


class SentimentType(Enum):
    HIGH = 3
    MIDDLE = 2
    LOW = 1


class FieldConstants(object):
    # 企业名称
    companyName = ""

    # 标题
    title = ""

    # 分类, SentimentType枚举
    Atype = SentimentType.MIDDLE

    # 阅读量
    readNum = 0

    # 评论量
    reviewNum = 0

    # 转发量
    retweetNum = 0

    # 点赞量
    likeNum = 0

    # 任务创建时间
    createTime = datetime.utcnow()

    # 文章发布时间
    _publishTime = datetime

    def __getattribute__(self, item):
        if item in ["likeNum", "retweetNum", "reviewNum", "readNum"]:
            return self.judge_int(item)
        if item == "Atype":
            return self.judge_enum(item)
        return object.__getattribute__(self, item)

    def judge_enum(self, name):
        value = self.__dict__.get(name)
        _type = type(value)
        if value is None:
            msg = NOVALUE.format(key=name, content=self.__dict__)
            raise ValueError(msg)
        if _type is SentimentType:
            return value.value
        else:
            msg = NOVALUE.format(key=name, content=self.__dict__)
            raise ValueError(msg)

    def judge_int(self, name):
        value = self.__dict__.get(name)
        if value is None:
            value = 0
        _type = type(value)

        if _type is int:
            return value
        else:
            msg = TPYEERROR.format(key=name, Type=int, content=self.__dict__)
            raise Exception(msg)

    @property
    def publish_time(self):
        return self._publishTime

    @publish_time.setter
    def publish_time(self, value):
        # 预处理
        if int == type(value):
            value = str(value)

        if str == type(value):
            if value.startswith("15") and (len(value) == 10 or len(value) == 13):
                value = datetime.fromtimestamp(int(value[:-3]))
            elif "20" in value:
                value = parse(value)

        # 判断
        if type(value) == datetime:
            self._publishTime = value
        else:
            msg = TPYEERROR.format(key="publish_time", Type=datetime, content=self.__dict__)
            raise ValueError(msg)

    def get_dict(self):
        pr = {}
        for name in dir(self):
            value = getattr(self, name)

            if name in ["companyName", "title", "createTime", "publish_time"]:
                if not value is None:
                    continue
                msg = NONEVALUE.format(name)
                raise ValueError(msg)

            if not name.startswith('__') and not callable(value) and not name.startswith('_'):
                pr[name] = value
        return pr

if __name__ == '__main__':
    s = FieldConstants()
    s.Atype = SentimentType.MIDDLE
    print(s.Atype)
