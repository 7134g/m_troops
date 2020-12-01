"""
外观模式（Facade），亦称“过程模式”， 为子系统中的一组接口提供一个一致的界面

Facade模式定义了一个高层接口，这个接口使得这一子系统更加容易使用。
与接口相关的适配器模式《python设计模式（五）：适配器模式——各种类接口的合并》有所不同的是外观模式是为大系统下的小系统设计统一的接口
而适配器模式是针对不同系统各种接口调用而设计。

在以下情况下可以考虑使用外观模式：

(1)设计初期阶段，应该有意识的将不同层分离，层与层之间建立外观模式。

(2) 开发阶段，子系统越来越复杂，增加外观模式提供一个简单的调用接口。

(3) 维护一个大型遗留系统的时候，可能这个系统已经非常难以维护和扩展，但又包含非常重要的功能，为其开发一个外观类，以便新系统与其交互。 [

优点
（1）实现了子系统与客户端之间的松耦合关系。

（2）客户端屏蔽了子系统组件，减少了客户端所需处理的对象数目，并使得子系统使用起来更加容易。
"""

class Facade:
    # def __init__(self):
    #     self.app1 = App1()
    #     self.app2 = App2()

    def postAll(self):
        print("post all")
        [obj.post() for k, obj in self.__dict__.items()]


    def delete(self):
        print("delete all")
        [obj.delete() for k, obj in self.__dict__.items()]

    def __getattr__(self, item):
        return getattr(self.name, item)


class App1:
    def post(self):
        print("app1 post")

    def delete(self):
        print("app1 delete")

class App2:
    def post(self):
        print("app2 post")

    def delete(self):
        print("app2 delete")

if __name__ == '__main__':
    facade = Facade()
    facade.app1 = App1()
    facade.app2 = App2()
    facade.postAll()
