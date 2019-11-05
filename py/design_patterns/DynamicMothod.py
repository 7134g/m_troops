"""
装饰模式指的是在不必改变原类文件和使用继承的情况下，动态地扩展一个对象的功能。它是通过创建一个包装对象，也就是装饰来包裹真实的对象。
用于动态添加属性或者方法 动态继承
适合用于元类编程
"""

class Dynamic:
    def __init__(self, classname):
        self._mothod = classname

    def __getattr__(self, item):
        return getattr(self._mothod, item)

    def test(self):
        print("dynamic test")

class Action:
    def test(self):
        print("action test")

    def singing(self):
        print("I can singing")

    def jump(self):
        print("I can jump")

    def rap(self):
        print("I can rap")

    def basketball(self):
        print("I can basketball")

def xswl():
    print("笑死我了")

# 动态赋值
dynamic = Dynamic(Action())
dynamic.name = "cxk"
dynamic.xswl = xswl

# 方法相同时执行Dynamic类中的方法
dynamic.test()

"""just do it"""
print(dynamic.name)
dynamic.singing()
dynamic.jump()
dynamic.rap()
dynamic.basketball()
dynamic.xswl()

