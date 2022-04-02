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

    def execute(self, data):
        return ""


def __init__(self):
    self.childrens = {}


def run(self, data, *args):
    for key, value in self.childrens.items():
        print('开始执行: {}'.format(key))
        result = value.execute(data)
        print(result)


def parent(self, *args):
    for index, value in enumerate(args):
        key = 'ChlidSystem{index}'.format(index=index)
        obj = value()
        self.childrens[key] = Adapter(obj, dict(execute=obj.do_something))


def main2():
    classname = "Adapter2"
    classtype = (object,)
    classdict = {
        "__init__": __init__,
        "parent": parent,
        "run_for": run,
    }
    Deom = type(classname, classtype, classdict)
    test = Deom()
    test.parent(S1, N1) # 接受不限量接口
    test.run_for("   main2")


if __name__ == '__main__':
    main2()