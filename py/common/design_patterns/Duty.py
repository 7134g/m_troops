"""
责任链模式
这条链条是一个对象包含对另一个对象的引用而形成链条，每个节点有对请求的条件，当不满足条件将传递给下一个节点处理。
"""


class Bases:

    def __init__(self, obj=None):
        self.obj = obj

    def screen(self, number):
        pass


class Top(Bases):

    def screen(self, number):

        if 200 > number > 100:
            print("{} 划入A集合".format(number))
        else:
            self.obj.screen(number)


class Second(Bases):

    def screen(self, number):

        if number >= 200:
            print("{} 划入B集合".format(number))
        else:
            self.obj.screen(number)


class Third(Bases):

    def screen(self, number):

        if 100 >= number:
            print("{} 划入C集合".format(number))


if __name__ == '__main__':

    test = [10, 100, 150, 200, 300]
    c = Third()
    b = Second(c)
    a = Top(b)
    for i in test:
        a.screen(i)
