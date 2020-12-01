"""
树结构

- 总经理办公室
    |- 财务部门
    |- 业务部门
        |- 销售一组
        |- 销售二组
    |- 生产部门
        |- 研发组
        |- 测试组
"""

class Node:

    def __init__(self, name, duty):
        self.name = name
        self.duty = duty
        self.children = []

    def add(self, obj):
        self.children.append(obj)

    def remove(self, obj):
        self.children.remove(obj)

    # 递归打印
    def display(self, number=1):
        print("{}部门：{} 层级：{} 职责：{}".format((number-1)*"\t",self.name, number, self.duty))

        n = number+1
        for obj in self.children:
            obj.display(n)


if __name__ == '__main__':
    # 顶级
    root = Node("总经理办公室", "总负责人")
    # 二级
    money = Node("财务部门", "公司财务管理")
    operation = Node("业务部门", "销售产品")
    production = Node("生产部门", "生产产品")
    root.add(money)
    root.add(operation)
    root.add(production)
    # 三级
    sell_first = Node("销售一组", "A产品销售")
    sell_second = Node("销售二组", "B产品销售")
    operation.add(sell_first)
    operation.add(sell_second)
    # 三级
    creat = Node("研发组", "研发组组长")
    test = Node("测试组", "测试组组长")
    production.add(creat)
    production.add(test)

    root.display()

