"""
flyweight pattern 享元模式
系统中存在大量的相似对象时，可以选择享元模式提高资源利用率。
"""

class FlyweightBase:

    def offer(self):
        """享元基类"""
        pass


class Flyweight(FlyweightBase):
    """共享享元类"""
    def __init__(self, name):
        self.name = name

    def get_price(self, price):
        print('产品类型：{} 详情：{}'.format(self.name, price))


class FactoryFlyweight:
    """享元工厂类"""
    def __init__(self):
        self.product = {}

    def Getproduct(self, key):
        if not self.product.get(key, None):  # 若键存在，则继续使用旧有的对象，操作同一片内存空间
            self.product[key] = Flyweight(key)
        return self.product[key]


if __name__ == '__main__':
    test = FactoryFlyweight()

    A = test.Getproduct("高端")
    A.get_price("香水：80")

    # 此时B使用的是A中的Flyweight对象
    B = test.Getproduct("高端")
    B.get_price("面膜：800")