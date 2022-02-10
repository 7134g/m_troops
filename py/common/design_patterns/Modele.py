"""
模块方法模式，一种装载的方式
实现了动态的更新流程或算法
"""

"""实现一个客户点单后的处理流程流程"""


class User:

    def __init__(self, name, shop, times, number):
        self.name = name
        self.shop = shop
        self.times = times
        self.number = number


class Handle:
    def __init__(self, user=None):
        self.user = user

    def invoicen(self):
        """打印小票"""
        string = "打印小票" \
                 "客户：{}" \
                 "商品：{}" \
                 "数量：{}" \
                 "时间：{}".format(self.user.code, self.user.shop, self.user.number, self.user.times)
        print(string)

    def make(self):
        """开始制作"""
        print("制作完成：{} 数量：{}".format(self.user.shop, self.user.number))

    def run(self):
        self.invoicen()
        self.make()


if __name__ == '__main__':
    test = Handle()

    xiaoming = User("小明", "汉堡", "17:50", "5")
    test.user = xiaoming
    test.run()

    xiaohong = User("小红", "北京卷", "18:00", "2")
    test.user = xiaohong
    test.run()
