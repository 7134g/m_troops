"""
观察者模式
动态监控值的变化，执行需要的操作
"""


class Observer:
    """观察者核心：销售人员，被观察者number数据"""
    def __init__(self):
        self._number = None
        self._department = []

    @property
    def number(self):
        return self._number

    # 每一次被赋值将会自动调用一次
    @number.setter
    def number(self, value):
        self._number = value
        print('当前客户数：{}'.format(self._number))
        for obj in self._department:
            obj.change(value)
        print('------------------')

    def notice(self, *args):
        """相关部门"""
        for d in args:
            self._department.append(d)


class Hr:
    """人事部门"""
    def change(self, value):
        if value < 10:
            print("人事变动：裁员")

        elif value > 20:
            print("人事变动：扩员")

        else:
            print("人事不受影响")


class Factory:
    """工厂类"""
    def change(self, value):
        if value < 15:
            print("生产计划变动：减产")
        elif value > 25:
            print("生产计划变动：增产")
        else:
            print("生产计划保持不变")


def alter(observer):
    observer.number = 10
    observer.number = 15
    observer.number = 20
    observer.number = 25


if __name__ == '__main__':
    observer = Observer()
    # 需要监控的对象
    hr = Hr()
    factory = Factory()
    # 装载监控对象
    observer.notice(hr,factory)
    # 数值变化
    alter(observer)
