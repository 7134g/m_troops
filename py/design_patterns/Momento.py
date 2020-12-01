"""
备忘录:
在不破坏封闭的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态。
这样以后就可将该对象恢复到原先保存的状态。
简单来说在运行过程中我们可以记录某个状态，当遇到错误时恢复当前状态，这在业务流程中是用设计来处理异常情况。
"""


import copy


class AddNumber:

    def __init__(self):
        self.start = 1

    def add(self, number):
        self.start += number
        print("传入的值 {},此时start的值 {}".format(number, self.start))


class Memento:
    """备忘录"""
    def backups(self, obj=None):
        """
        设置备份方法
        """
        self.obj_dict = copy.deepcopy(obj.__dict__)
        print("备份数据:{}".format(self.obj_dict))

    def recovery(self, obj):
        """
        恢复备份方法
        """
        obj.__dict__.clear()
        obj.__dict__.update(self.obj_dict)
        return obj


if __name__ == '__main__':
    test = AddNumber()
    memento = Memento()

    for i in [1, 2, 3, 'n', 4]:
        if i == 2:
            # 记录保存test栈信息
            memento.backups(test)
        try:
            test.add(i)
        except TypeError as e:
            print("传入的值 {}，此时start的值 {}, 发生错误信息如下： {}".format(i, test.start, e))

    # 还原之前保存的数据
    memento.recovery(test)
    # 查看保存的数据值
    print("还原数据，此时start的值 {}".format(test.start))
