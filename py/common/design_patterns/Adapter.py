# 接口1
class S1:
    def do_something(self, data):
        return "".join(["S1", data])
# 接口2
class N1:
    def do_something(self, data):
        return "".join(["N1", data])

# 适配器
class Adapter:
    def __init__(self, obj, adapted_methods):
        self.obj = obj
        self.__dict__.update(adapted_methods)

    def __str__(self):
        return str(self.obj)


def deal_msg():
    # 生成实例
    objects = {}
    s = S1()
    n = N1()
    # 注册
    objects["1"] = Adapter(s, dict(execute=s.do_something))
    objects["2"] = Adapter(n, dict(execute=n.do_something))

    return objects


if __name__ == '__main__':
    config = "1"
    data = "data"
    s = deal_msg()[config].execute(data)
    print(s)