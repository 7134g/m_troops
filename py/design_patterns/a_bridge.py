# 桥接模式
# 与适配器模式区别在于 只容纳一个类 // 适配器可容纳多个类

class Bridge:
    def __init__(self, class_example, data):
        self.class_example = class_example
        self.data = data

    def bridge_run(self):
        return self.class_example.run(self.data)

class Chinese:
    def run(self,data):
        '''deal something'''
        print("中文：{}".format(data))
        return True

class English:
    def run(self,data):
        '''deal something'''
        print("English：{}".format(data))
        return True

if __name__ == '__main__':
    b = Bridge(Chinese(), "世界你好")
    b.bridge_run()

    b = Bridge(English(), "hellow world")
    b.bridge_run()