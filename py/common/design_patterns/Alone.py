import threading
# 单例模式 ,无论实例化多少次，都是操作同一片内存空间

# 方法一
class Singleton(object):

    def __new__(cls, *args, **kw):
        if not hasattr(cls, '_instance'):
            org = super(Singleton, cls)
            cls._instance = org.__new__(cls, *args, **kw)
        return cls._instance

    def run(self, count=1):
        if count != 1:
            self.count = count
        print(self.count)

# 方法二，元类
class SingletonType(type):
    _instance_lock = threading.Lock()
    def __call__(cls, *args, **kwargs):
        if not hasattr(cls, "_instance"):
            with SingletonType._instance_lock:
                if not hasattr(cls, "_instance"):
                    cls._instance = super(SingletonType,cls).__call__(*args, **kwargs)
        return cls._instance

class Foo(metaclass=SingletonType):
    def __init__(self,name):
        self.name = name

if __name__ == '__main__':
    s = Singleton()
    s.run(count=100)

    e = Singleton()
    e.run()

    print(id(s), id(e))
    print(s==e)

    print("================================")
    f1 = Foo("liming")
    print(f1.name)
    f2 = Foo()
    print(f2.name)
    print(id(f1), id(f2))