#### 安装服务器
    yum install mosquitto -y
#### 运行服务器
    mosquitto

#### 分布式 emqttd
    yum -y install make gcc gcc-c++ kernel-devel m4 ncurses-devel openssl-devel
    wget http://www.erlang.org/download/otp_src_R13B04.tar.gz
    tar xfvz otp_src_R13B04.tar.gz
    cd otp_src_R13B04/
    ./configure --with-ssl
    sudo make install
    来源：
    https://blog.csdn.net/laughing_cui/article/details/53322790
    