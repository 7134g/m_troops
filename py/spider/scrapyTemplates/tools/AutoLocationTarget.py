# -!- coding: utf-8 -!-
import re

from bs4 import BeautifulSoup
from scrapy.http.cookies import CookieJar
from parsel import Selector
from py.spider.scrapyTemplates.tools.DataStruct import ResultInfo


# 获取response中返回的cookies
def get_cookies(response) -> dict:
    cookie_dict = {}
    cookie_jar = CookieJar()
    cookie_jar.extract_cookies(response, response.request)
    for k, v in cookie_jar._cookies.items():
        for i, j in v.items():
            for m, n in j.items():
                cookie_dict[m] = n.value
    return cookie_dict


# 清理所有html标签和空白字符
def delete_lable(_html: str = "", _pattern: str = "") -> str:
    if not _html:
        return ""
    parting = r"</?\w+[^>]*>" + _pattern
    return re.sub(parting, "", _html)


# 清理所有html标签和空白字符
def delete_lable_and_space(_html: str = "", _pattern: str = "") -> str:
    if not _html:
        return ""
    parting = r"</?\w+[^>]*>|\s" + _pattern
    return re.sub(parting, "", _html)


# 通过某一文本内容定位至目标位置，然后逐级向外寻找相同的父标签(xpath)
def xpath_local_content(text, local_point="电话", contrast_point="地址") -> list:
    content = text.replace("&#12288;", "").replace("\u3000", "")
    # content = response.replace("&#12288;", "").replace("\u3000", "")
    selete = Selector(content)
    targets = [
        '//*[contains(text(),"{}")]'.format(local_point),
        '//*[contains(text(),"手机")]',
        '//*[contains(text(),"座机")]']
    second_target = '//*[contains(text(),"{}")]'.format(contrast_point)
    lenght = len(targets)
    first_local = []
    index = 0

    for index, first_value in enumerate(targets):
        first_local = selete.xpath(first_value)
        if first_local:
            break
    if not first_local:
        return []

    while index + 1 <= lenght:
        for count in range(1, 4):
            find_parent = '/..' * count
            first_local = selete.xpath(targets[index] + find_parent)
            second_local = selete.xpath(second_target + find_parent).extract()
            for second_value in second_local:
                for offset, value in enumerate(first_local):
                    result = value.extract()
                    # print(result)
                    if result == second_value:
                        result = selete.xpath(targets[index] + find_parent + '/*').extract()
                        return result
        index += 1

    return []


# 通过某一文本内容定位至目标位置，然后逐级向外寻找相同的父标签(bs4)
def bs4_local_content(text, local_point="电话", *args, contrast_point="地址", single=1):
    soup = BeautifulSoup(text, "lxml")

    all_search = [local_point, *args, "手机", "座机", "电话"]
    for each in all_search:
        arm = '.*{}.*'.format(each)
        arm2 = '.*{}.*'.format(contrast_point)

        reg = re.compile(arm)
        reg2 = re.compile(arm2)

        first_point = soup.find_all(text=reg)
        second_point = soup.find_all(text=reg2)

        # 假如层级不一致时
        for fast in range(3):
            # 假如父节点并不是同一个，需要父节点的父节点
            for count in range(1, 4):
                tag1_exec_str = "f" + ".parent" * count
                tag2_exec_str = "s" + ".parent" * (count + fast)
                # 两个定位点对比
                for f in first_point:
                    for s in second_point:
                        tag1 = eval(tag1_exec_str)
                        tag2 = eval(tag2_exec_str)
                        if tag1 == tag2:
                            # 成功得出结果
                            # 纯字符串
                            if single == 1:
                                result = tag1.stripped_strings
                                return list(result)

                            # 保留标签
                            if single == 2:
                                result = tag1.strings
                                return list(result)
    return ""


# 根据映射表给 ResultInfo 对象赋值
def exec_map_contract(map_table: dict, target, data: ResultInfo):
    _type = type(target)
    for each in target:
        if _type == list:
            p_data = delete_lable_and_space(each).split("：")
            search_arm = map_table.get(p_data[0], None)
            if search_arm:
                exec_str = "data.{} = '{}'".format(search_arm, p_data[1])
                exec(exec_str)
        else:
            # p_data = delete_lable_and_space(each)
            key = each
            search_arm = map_table.get(key, "")
            if search_arm:
                exec_str = "data.{} = '{}'".format(search_arm, target[key])
                exec(exec_str)



def get_bs4map_table(target: list, clear=False, **kwargs) -> dict:
    results = target
    if not results:
        return {}
    count = len(results)
    new_result = {}
    # try:
    for each in range(count):
        content = results[each]
        if clear:
            content = delete_lable_and_space(content, **kwargs)

        # 判断是否有"："， 如果有分割后左为键，右为值
        if "：" in content or ':' in content:
            try:
                # 下一个元素没有冒号
                if "：" not in results[each + 1] or ':' not in results[each + 1]:
                    # 例子：['地址：广东省广州市', '无用数据']
                    # 后面无数据
                    if not (content.find("：") + 1 or content.find(":") + 1) == len(content):
                        result = content.split("：")
                        if len(result) == 1:
                            result = content.split(":")
                        key = result[0]
                        try:
                            value = result[1]
                        except IndexError:
                            value = ""
                    else:
                        # 例子：['地址：', '广东省广州市']
                        # 后面为前面的值
                        key = content.replace("：", "").replace(":", '')
                        value = results[each + 1]

                # 下一个元素有冒号
                # 例子：['地址：广东省广州市', '电话：16655484894']
                else:
                    result = content.split("：")
                    if len(result) == 1:
                        result = content.split(":")
                    key = result[0]
                    try:
                        value = result[1]
                    except IndexError:
                        value = ""
            except IndexError:
                result = content.split("：")
                if len(result) > 1:
                    key = result[0]
                    value = result[1]
                else:
                    key = content
                    value = ""
            new_result.update({key: value})

    return new_result



def test(html, xp=True, bs=False):
    map_table = {
        "手机": "phone",
        "移动电话": "phone",
        "公司电话": "landline",
        "电话": "landline",
        "公司传真": "fax",
        "传真": "fax",
        "地址": "address",
        "公司地址": "address",
        "详细地址": "address",
        "邮箱": "mail",
        "E-Mail": "mail",
        "微信号码": "wechat",
        "联系人": 'contactPerson',
        "E-mail": "mail",
        "腾讯QQ": "qq",
        '公司名称': "companyName",
    }

    """feijiu  hxsteel  atobo  sonhoo  agronet TFE yhby"""

    if bs:
        # 用例 一  bs4_local_content
        data = ResultInfo()
        result = bs4_local_content(html, local_point="联系人", contrast_point="公司主页")
        # result = eval(re.sub(r'◆|\\xa0|\s', "", str(result)))
        # result = eval(delete_lable_and_space(str(result), _pattern="\\◆|\\\\xa0"))
        # print(result)
        # print(type(result))
        result_map = get_bs4map_table(result)
        exec_map_contract(map_table, result_map, data)

        print(result_map)
        print(vars(data))
        return data

    if xp:
        # 用例 三  xpath_local_content
        data = ResultInfo()
        result = xpath_local_content(html, local_point="电话")
        print(type(result))
        exec_map_contract(map_table, result, data)
        print(vars(data))
        return data


if __name__ == '__main__':
    html = "..."
    test(html)