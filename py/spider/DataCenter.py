import re
from datetime import datetime, timedelta
from dateparser import parse as d_parse
from dateutil.parser import parse


def is_true(value):
    judge_list = [False, " "]
    if value not in judge_list or isinstance(value, int):
        return value


def json_to_empty(value, layer=1):
    if layer>10:
        raise IndexError("超出最大递归层数")
    if type(value) == list:
        new_list = []
        for i in value:
            _layer = layer + 1
            new_value = json_to_empty(i, _layer)
            judge = is_true(new_value)
            if judge or isinstance(judge, int):
                new_list.append(new_value)
        if new_list:
            return new_list
    elif type(value) == dict:
        new_dict = {}
        for k, v in value.items():
            judge = is_true(v)
            if judge or isinstance(judge, int):
                _layer = layer + 1
                new_value = json_to_empty(judge, _layer)
                judge = is_true(new_value)
                if judge or isinstance(judge, int):
                    new_dict[k] = new_value
        return new_dict
    else:
        judge = is_true(value)
        if judge or isinstance(judge, int):
            return judge


class JudgeField:
    NUMTYPE = [""]
    STRTYPE = ["mainBusiness",
               "creditCode",
               "address",
               "area",
               "businessScope",
               "position",
               "mail",
               "wechat",
               "qq",
               "keywords",
               "industry",
               "contactPerson",
               "type",
               "sourceUrl",
               "companyWebSite"]
    NONETYPE = ["companyName", "source", "contactUrl"]
    TIMETPYE = ["publishTime", ]
    ENUMTPYE = ["category"]
    PHONE = ["phone", "landline"]
    PACKTYPE = ["financing", "brand", "competing_goods", "company_news", "icp"]
    COGNIZETYPE = ["cognizanceType"]  # 认证类型


class StructBase:
    def __getattribute__(self, item):
        if item in JudgeField.NUMTYPE:
            return self.deal_int(item)  # 数字
        elif item in JudgeField.STRTYPE:
            return self.deal_str(item)  # 字符串
        elif item in JudgeField.NONETYPE:
            return self.deal_str(item)  # 空值
        # elif item in SalesJudgeField.ENUMTPYE:
        #     return self.deal_enum(item)  # 枚举
        return object.__getattribute__(self, item)

    def __setattr__(self, key, value):
        if key in JudgeField.TIMETPYE:
            return self.deal_datetime(key, value)
        elif key in JudgeField.PHONE:
            return self.deal_phone(key, value)
        elif key in JudgeField.PACKTYPE:
            return self.deal_pack(key, value)
        object.__setattr__(self, key, value)

    def deal_pack(self, key, value):
        now_value = self.__dict__.get(key, [{}])
        if now_value == [{}]:
            self.__dict__[key] = []
            self.__dict__[key].append(value)
            return

        # 去重
        compare_max = len(value)
        num = 0
        for each in now_value:  # 获取现有的所有dict
            for k in value:  # 将每一个键值对比
                if value[k] == each[k]:
                    num += 1
            if num == compare_max:  # 若全等，则重复
                return

            num = 0  # 下一个dict重置计数器


        # 赋值
        self.__dict__[key].append(value)

    def deal_str(self, item):
        value = self.__dict__.get(item)
        if not value:
            return None
        result = re.sub('\s', "", value)
        return result

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
            print(f"{key} 无法parse成 datetime 对象, 来源: {self.companyName}")
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
            landline_turn = ''
            if not value:
                return

            if len(value) < 5:
                return

            # 清洗
            v_split = re.split('\D', value)
            if len(v_split) != 1:  # 判断是否全数字
                # 判断是否可能为phone
                for index, each in enumerate(v_split):  # 遍历分割后每段是否可能是手机号
                    net_id = each[:3]
                    if net_id in phone_three:
                        key = "phone"
                        value = ''.join(v_split[index:])  # 合并后面全部数据
                        # print(value)
                        break

                if key == "phone":  # 判断类型
                    _v = re.sub(r"\D", "", value)
                    if len(_v) != 11:  # 字符串数字不是11位长度
                        return
                    _v = re.findall(r'\d{11}', _v)
                    _v = _v[0] if _v else ""

                elif key == "landline":
                    _v = re.findall(r'\d+.\d+.\d+|86.?\d+.\d+', value, re.S)
                    landline_turn = re.findall('转\d+', value)
                    landline_turn = landline_turn[0] if landline_turn else ''
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

            # 手机号前混杂其他数字
            for three in phone_three:
                if three in value:
                    temp = value[value.find(three):]
                    if len(temp) == 11:
                        _v = temp
                        break

            _v1 = re.sub(r"\D", "", _v)

            # 去重
            for each in JudgeField.PHONE:
                is_exist = self.__dict__.get(each, [])
                if not is_exist:
                    self.__dict__[each] = []

                if _v in is_exist or _v1 in is_exist:
                    return

            if len(_v1) < 5 or len(_v) < 5:
                return

            if _v1[:3] in phone_three:
                self.__dict__["phone"].append(_v1)
            else:
                # 分机号
                if landline_turn:
                    _v = re.sub(r"\D", "", _v)
                    _v += landline_turn.replace('转', '-')
                else:
                    _v = re.sub(r"\D", "-", _v)
                self.__dict__["landline"].append(_v)

    def get_dict(self):
        pr = {}
        for name in dir(self):
            value = getattr(self, name)

            if not name.startswith('__') and not callable(value) and not name.startswith('_'):
                if value is None:
                    msg = "{} 的值为空, 来源于: {}".format(name, getattr(self, "eid"))
                    raise ValueError(msg)
                pr[name] = value
        new_data = json_to_empty(pr)
        return new_data


if __name__ == '__main__':
    s = StructBase()
    s.publishTime = datetime.now()
    s.companyName = "111"
    s.title = "222"
    print(s.get_dict())
