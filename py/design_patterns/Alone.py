# 单例模式 ,无论实例化多少次，都是操作同一片内存空间

class Singleton(object):

    def __new__(cls, *args, **kw):
        if not hasattr(cls, '_instance'):
            org = super(Singleton, cls)
            cls._instance = org.__new__(cls, *args, **kw)
        return cls._instance

    def run(self):
        print("run")

if __name__ == '__main__':
    s = Singleton()
    s.run()

    e = Singleton()
    e.run()

    print(id(s), id(e))
    print(s==e)