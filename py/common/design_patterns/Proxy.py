"""
远程代理：
    为一个位于不同的地址空间的对象提供一个局域代表对象。这个不同的地址空间可以是在本机器中，也可是在另一台机器中。
    远程代理又叫做大使（Ambassador）。好处是系统可以将网络的细节隐藏起来，使得客户端不必考虑网络的存在。

虚拟代理（图片延迟加载的例子）：
    根据需要创建一个资源消耗较大的对象，使得此对象只在需要时才会被真正创建。
    使用虚拟代理模式的好处就是代理对象可以在必要的时候才将被代理的对象加载；代理可以对加载的过程加以必要的优化。
    当一个模块的加载十分耗费资源的情况下，虚拟代理的好处就非常明显。

保护代理：
    控制对原始对象的访问。保护代理用于对象应该有不同 的访问权限的时候

智能引用代理：
    当一个对象被引用时，提供一些额外的操作，比如将对此对象调用的次数记录下来等
"""

class Jurisdiction:
    """权限类"""

    def level1(self):
        print('权限等级1')

    def level2(self):
        print('权限等级2')

    def level3(self):
        print('权限等级3')

    def level4(self):
        print('权限等级4')


class Proxy:

    def __init__(self, name):
        self.user = name
        self._jurisdiction = Jurisdiction()

    def leve(self):

        if self.user == 'a':
            return self._jurisdiction.level1()
        elif self.user == 'b':
            return self._jurisdiction.level2()
        elif self.user == 'c':
            return self._jurisdiction.level3()
        elif self.user == 'd':
            return self._jurisdiction.level4()
        else:
            print('无此权限')


if __name__ == '__main__':
    test = Proxy('a')
    test.leve()
    test.user = 'b'
    test.leve()
    test.user = 'c'
    test.leve()
    test.user = 'd'
    test.leve()
    test.user = 'e'
    test.leve()
