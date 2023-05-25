##### 通过注册表创建桌面图标 以及 添加删除程序中显示卸载按钮
1. 创建文件夹
```
计算机\HKEY_LOCAL_MACHINE\SOFTWARE\WOW6432Node\Microsoft\Windows\CurrentVersion\Uninstall\Project
```
2. 注册表变量描述
```
EstimatedSize   描述大小（16进制）  6747e
DisplayIcon     描述图标            C:\project_name\Project.exe
DisplayName     描述名              Project
Publisher       发布者              compare
URLInfoAbout    发布者url地址       https://www.test.com/
UninstallString 卸载程序路径        C:\project_name\uninstall.exe
```

#### 提权
```
// icacls C:\Windows\System32\wpcap.dll /grant Users:(F)
mod := fmt.Sprintf(`icacls %s /grant Users:(F)`, dst)
```

#### 给予管理员权限
- 方法一
```
1> go get github.com/akavel/rsrc
2> 把nac.manifest 文件拷贝到当前windows项目根目录
3> rsrc -manifest nac.manifest -o nac.syso
或者
   rsrc -manifest nac.manifest -o nac.syso -arch amd64 
4> go build


<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<assembly xmlns="urn:schemas-microsoft-com:asm.v1" manifestVersion="1.0">
<trustInfo xmlns="urn:schemas-microsoft-com:asm.v3">
<security>
<requestedPrivileges>
<requestedExecutionLevel level="requireAdministrator"/>
</requestedPrivileges>
</security>
</trustInfo>
</assembly>

// 出现gcc问题
go build -ldflags "-linkmode internal"
https://blog.csdn.net/wn0112/article/details/106512945/
```
- 方法二
```
https://github.com/tc-hib/go-winres

// 下载
go install github.com/tc-hib/go-winres@latest
// 初始化
go-winres init
// 赋予管理员权限
修改winres/winres.json中的execution-level值为"administrator"
// 生成syso文件
go-winres make
```


#### 重新安装服务(管理员权限)
- 停止服务
    - cmd代码：sc stop npf
- 删除服务
    - cmd代码：sc delete npf
- 删除关于服务的注册表内容
    - go代码：registry.DeleteKey
- 删除驱动文件
    - go代码：os.Remove
- 创建新的驱动文件
    - go代码：CopyFile
- 添加注册表服务信息
    - go代码：registry.CreateKey
- 创建服务
    - cmd代码：sc create npf binPath= "system32\drivers\NPF.sys" type= kernel start= auto DisplayName= "WinPcap Packet Driver (NPF)"
- 启动服务
    - cmd代码：sc start npf


```
创建服务命令 sc create ...... 用法
例子一：
sc create npf binPath= "system32\drivers\NPF.sys" type= kernel start= auto DisplayName= "WinPcap Packet Driver (NPF)"

SERVICE_NAME: npf
        TYPE               : 1  KERNEL_DRIVER
        START_TYPE         : 2   AUTO_START
        ERROR_CONTROL      : 1   NORMAL
        BINARY_PATH_NAME   : system32\drivers\NPF.sys
        LOAD_ORDER_GROUP   :
        TAG                : 0
        DISPLAY_NAME       : WinPcap Packet Driver (NPF)
        DEPENDENCIES       :
        SERVICE_START_NAME :


例子二：
sc create npcap binPath= "\SystemRoot\system32\DRIVERS\npcap.sys" type= kernel start= system group= NDIS tag= yes DisplayName= "Npcap Packet Driver (NPCAP)"

SERVICE_NAME: npcap
        TYPE               : 10  WIN32_OWN_PROCESS
        START_TYPE         : 3   DEMAND_START
        ERROR_CONTROL      : 1   NORMAL
        BINARY_PATH_NAME   : C:\Windows\System32\drivers\npcap.sys
        LOAD_ORDER_GROUP   :
        TAG                : 0
        DISPLAY_NAME       : npcap
        DEPENDENCIES       :
        SERVICE_START_NAME : LocalSystem


选项:

注意: 选项名称包括等号。

      等号和值之间需要一个空格。
      
type= <own|share|interact|kernel|filesys|rec>

       (默认 = own)

 start= <boot|system|auto|demand|disabled|delayed-auto>

       (默认 = demand)

 error= <normal|severe|critical|ignore>

       (默认 = normal)

 binPath= <BinaryPathName>

 group= <LoadOrderGroup>

 tag= <yes|no>

 depend= <依存关系(以 / (斜杠) 分隔)>

 obj= <AccountName|ObjectName>

       (默认 = LocalSystem)

 DisplayName= <显示名称>

  password= <密码>
  
 
```

#### 添加到注册表
- [go code](registry_windows.go)

#### 创建服务
- [go code](serve_windows.go)

#### 修改cmd窗口标题
- [go code](cmd_windows.go)

#### 生成快捷方式
- [go code](link_windows.go)

#### 硬盘操作
- [go code](dist_windows.go)