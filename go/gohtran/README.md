Golang版本的Htran工具

#### file explain
    - main.go 命令解析及执行
    - config.go 全局变量配置信息
    - design.go 三种连接方式
    - connect.go 处理所有连接
    - fd.go 操作传输文件数据
    - aes.go aes相关操作
    - gzip.go gzip相关操作
    - util.go 
    - model.go 模块

#### 功能介绍
- 本地监听转发
    - gohtran -listen 本地端口1 本地端口2  
        ```gohtran -listen 1997 2017```
- 转发到远程主机
    - gohtran -tran 本地端口 远程ip:远程端口  
        ```gohtran -tran 2017 14.215.177.38:3389```
- 反向连接转发
    - gohtran -slave 指定主机端口 指定主机端口  
        ```gohtran -slave 127.0.0.1:2017 192.168.88.8:3389```
- aes加密解密  
    ```gohtran -listen 1997 2017 -aes```
- gzip压缩解压  
    ```gohtran -listen 1997 2017 -gzip```
- aes和gzip同时开启
    ```gohtran -listen 1997 2017 -e```
- 查看帮助  
    ```gohtran -h```  
    ```gohtran -help```  
- 静默模式
    ```gohtran -s```  
    ```gohtran -silent```  
   
#### 使用说明
    使用aes或者gzip，需要至少2个gohtran程序才能做到加密压缩操作
    - 例如
        本机 -  公网服务器  -  目标服务器
        tran - listen      - slave
        
#### 编译
    go build -ldflags "-s -w"
    2.5M