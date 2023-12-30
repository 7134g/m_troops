### websockets 优化

1. 提高操作系统的文件描述符限制

`ulimit -n 1000000` // 调整为百万连接

2. 优化tcp参数

`sysctl -w net.ipv4.tcp_fin_timeout=30` // 减少连接延迟

`sysctl -w net.ipv4.tcp_tw_reuse=1` // 资源消耗


