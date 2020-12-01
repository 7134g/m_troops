import re
from datetime import datetime, timedelta
from dateparser import parse as d_parse
from dateutil.parser import parse


class EumeOrigin(object):
    producers = "producers"
    Consumers = "Consumers"
    middlemen = "middlemen"


class JudgeField:
    NUMTYPE = ["count"]
    STRTYPE = ["content"]
    NONETYPE = ["task_name"]
    TIMETPYE = ["publish_time"]
    ENUMTPYE = ["category"]
    PHONE = ["phone", "landline"]


class ResultInfo:
    task_name = None  # 任务名
    content = ""  # 内容
    count = ""  # 数量
    publish_time = ""  # 发布时间
    category = ""  # 类型
    phone = ""  # 手机
    landline = ""  # 座机

    def __init__(self, spider_name=""):
        self._spiderName = spider_name

    def __getattribute__(self, item):
        if item in JudgeField.NUMTYPE:
            return self.deal_int(item)  # 数字
        elif item in JudgeField.STRTYPE:
            return self.deal_str(item, signal=1)  # 字符串
        elif item in JudgeField.NONETYPE:
            return self.deal_str(item, signal=2)  # 空值
        # elif item in SalesJudgeField.ENUMTPYE:
        #     return self.deal_enum(item)  # 枚举
        return object.__getattribute__(self, item)

    def __setattr__(self, key, value):
        if key in JudgeField.TIMETPYE:
            return self.deal_datetime(key, value)
        elif key in JudgeField.PHONE:
            return self.deal_phone(key, value)
        object.__setattr__(self, key, value)

    def deal_int(self, name):
        msg = "传入值有误, 此时数据包为: {}, {} 应该为 int ".format(self.__dict__, name)
        value = self.__dict__.get(name)
        if value is None:
            value = 0
        _type = type(value)

        if _type is int:
            return value
        elif _type is str:
            try:
                return int(value)
            except ValueError:
                pass
        else:
            raise ValueError(msg)

    def deal_str(self, name, signal):
        value = self.__dict__.get(name)
        if 1 == signal:
            if not value:
                return ""
        elif 2 == signal:
            if not value:
                return None
        try:
            result = re.sub(r"\s", "", value)
            result = result.replace("&nbsp", "")

            head_tail_removed_ch = [";", ":", " ", "：", "；", " ："]
            for ch in head_tail_removed_ch:
                result = result.strip(ch)
            return result
        except TypeError:
            print("ERROR TypeError: commons.SalesConstants.deal_str() 参数 {name}: {value} ".format(
                name=name, value=value))

    def deal_datetime(self, key, value):
        # 预处理
        if int is type(value) and (1900 < value < 2100 or value > 1000000000):
            value = str(value)

        if str is type(value) and value:
            if re.search(".*?([\u4E00-\u9FA5]+).*?", value):
                value = d_parse(value)
            else:
                value = parse(value)

            if value > datetime.now():
                value = value.replace(datetime.now().year - 1)

            value = value - timedelta(hours=8)

        # 判断
        if type(value) is datetime:
            self.__dict__[key] = value
        else:
            print(f"{key} 无法parse成 datetime 对象, 来源: {self.spiderName}")
            return ""

    def deal_phone(self, key, value):
        phone_three = [
            '130', '131', '132', '133', '134', '135', '136', '137',
            '138', '139', '145', '146', '147', '148', '149', '150',
            '151', '152', '153', '154', '155', '156', '157', '158',
            '159', '162', '165', '166', '167', '170', '170', '170',
            '171', '172', '173', '175', '176', '177', '178', '180',
            '181', '182', '183', '184', '185', '186', '187', '188',
            '189', '190', '191', '192', '193', '195', '196', '197',
            '198', '199']

        _type = type(value)

        # if not hasattr(self, key):
        #     self.__dict__[key] = []

        if _type == list:
            value = [each for each in value if each]
            self.__dict__[key] = value

        elif _type == str:
            if not value:
                return

            if len(value) < 5:
                return

                # 清洗
            if re.search(r'\D\d*\D', value):
                if key == "phone":
                    _v = re.sub(r"\D", "", value)
                    if len(_v) == 11:  # 字符串数字不是11位长度
                        return
                    _v = re.findall(r'\d{11}', _v)
                    _v = _v[0] if _v else ""

                elif key == "landline":
                    _v = re.findall(r'\d+.\d+.\d+|86.{0,1}\d+.\d+|\d+.\d+', value, re.S)
                    if _v:
                        _v = _v[-1]
                        if r"\n" in _v:
                            _v = _v.replace("\n", "-")
                    else:
                        _v = ""

                # 非 phone 和 landline
                else:
                    _v = value

            # 无需清洗
            else:
                _v = value

            # 没有值
            if not _v:
                return

            _v1 = re.sub(r"\s|-", "", _v)

            # 去重
            for each in JudgeField.PHONE:
                is_exist = self.__dict__.get(each, [])
                if not is_exist:
                    self.__dict__[each] = []

                if _v in is_exist or _v1 in is_exist:
                    return

            if len(_v1) < 5 or len(_v) < 5:
                return

            # 存储
            # 判断是否是手机号
            if _v1[:3] in phone_three:
                self.__dict__["phone"].append(_v1)
            else:
                self.__dict__["landline"].append(_v)

    def get_dict(self):
        pr = {}
        for name in dir(self):
            value = getattr(self, name)

            if not name.startswith('__') and not callable(value) and not name.startswith('_'):
                if value is None or isinstance(value, EumeOrigin):
                    msg = "{} 的值为空, 来源于: {}".format(name, getattr(self, "source"))
                    raise ValueError(msg)
                pr[name] = value
        return pr
