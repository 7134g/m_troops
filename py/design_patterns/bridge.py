class Bridge:
    def __init__(self, classname, data):
        self.classname = classname
        self.data = data

    def bridge_run(self):
        return self.classname.run(self.data)

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
    b = Bridge(Chinese(),"世界你好")
    b.bridge_run()

    b = Bridge(English(), "hellow world")
    b.bridge_run()