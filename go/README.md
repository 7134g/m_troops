## 安全指南
- [Tencent](https://github.com/Tencent/secguide/blob/main/Go%E5%AE%89%E5%85%A8%E6%8C%87%E5%8D%97.md)


## 可编译工具
- [tcp简易命令工具](project/tcpDialAndServe/README.md)
- [暴力破解压缩包](project/recursion_decode/README.md)
- [去除重复文件](project/duplication/README.md)
- [匿名邮件](project/stmp/README.md)
- [切割合并文件](project/split_merge_file/README.md)
- [广播](project/broadcast/main.go)
- [rpc简易命令工具](project/rpc/README.md)
- [etcd_client](project/rpc/README.md)
- [mqtt_client](project/mqtt/README.md)

## 代码块和库
- [常用库](common.md)
- [协程池](common/pool/pool.go)
- [web](web)
    - [gin](web/gin.md)
    - [go_zero](web/go_zero.md)
    - [go_micro](web/go_micro.md)
  
    - [validate](web/validate.md)
    - [swagger](web/swagger.md)
    - [nsq](web/nsq.md)
    - [gorm](common/db)
      - [mysql连接](common/db/mysql.go)
      - [redis连接](common/db/redis.go)
      - [redis分布式锁](common/db/redis_lock.go)
      - [mongo](common/db/mongo.go)
    - [自旋锁](common/cas/self_lock.go)
    - [编解码](common/encoding/coding.go)
    - [进制变化](common/encoding/convet.go)
    - [结构体转map](common/encoding/StructAssignment.go)
    - [文件处理](common/files/files.go)
    - [拦截http改包](common/proxy/martian.go)
    - [证书](common/proxy/cert.go)
    - [windows注册表及服务](common/system/windows.md)
- [夏普率](project/finance/sharpe.md)

    
## mock
  - [convey](common/mock_go/convey_test.go)
  - [mock](common/mock_go/mock_test.go)
  - [monkey](common/mock_go/monkey_test.go)
  - [monkey_doc](common/mock_go/gomonkey.md)
  - [sql_mock](common/mock_go/sql_test.go)


## 数据结构
  - [链表](common/data_structure/linked_list.go)

## 设计模式
- [抽象工厂模式](common/design_patterns/abstract_factory/abstractfactory.go)
- [适配器模式](common/design_patterns/adapter/adapter.go)
- [单例模式](common/design_patterns/alone/alone.go)
- [单例方法模式](common/design_patterns/alone_method/alone_method.go)
- [桥接模式](common/design_patterns/bridge/bridge.go)
- [生成器模式](common/design_patterns/builder/builder.go)
- [责任链模式](common/design_patterns/chain_of_responsibility/chain.go)
- [命令模式](common/design_patterns/command/command.go)
- [组合模式](common/design_patterns/composite/composite.go)
- [装饰模式](common/design_patterns/decorator/decorator.go)
- [外观模式](common/design_patterns/facade/facade.go)
- [工厂模式](common/design_patterns/factory/factory.go)
- [工厂方法模式](common/design_patterns/factory_method/factorymethod.go)
- [享元模式](common/design_patterns/flyweight/flyweight.go)
- [解释器模式](common/design_patterns/interpreter/interpreter.go)
- [迭代器模式](common/design_patterns/iterator/iterator.go)
- [中介者模式](common/design_patterns/mediator/mediator.go)
- [备忘录模式](common/design_patterns/memento/memento.go)
- [观察者模式](common/design_patterns/observer/observer.go)
- [原型模式](common/design_patterns/prototype/prototype.go)
- [代理模式](common/design_patterns/proxy/proxy.go)
- [状态模式](common/design_patterns/state/state.go)
- [策略模式](common/design_patterns/strategy/strategy.go)
- [模版方法模式](common/design_patterns/template_method/templatemethod.go)
- [访问者模式](common/design_patterns/visitor/visitor.go)


## 排序
[排序动态演示地址](https://visualgo.net/zh/sorting)

- [冒泡](common/data_structure/sort/bub.go)
- [归并](common/data_structure/sort/mer.go)
- [快排](common/data_structure/sort/qui.go)
