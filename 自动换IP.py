#!/usr/bin/env python
#-*- encoding:gb2312 -*-
# Filename: IP.py
import sitecustomize
import _winreg
import ConfigParser
from ctypes import *
print '正在进行网络适配器检测，请稍候…'
print
netCfgInstanceID = None
hkey = _winreg.OpenKey(_winreg.HKEY_LOCAL_MACHINE, \
r'System\CurrentControlSet\Control\Class\{4d36e972-e325-11ce-bfc1-08002be10318}')
keyInfo = _winreg.QueryInfoKey(hkey)
# 寻找网卡对应的适配器名称 netCfgInstanceID
for index in range(keyInfo[0]):
hSubKeyName = _winreg.EnumKey(hkey, index)
hSubKey = _winreg.OpenKey(hkey, hSubKeyName)
try:
hNdiInfKey = _winreg.OpenKey(hSubKey, r'Ndi\Interfaces')
lowerRange = _winreg.QueryValueEx(hNdiInfKey, 'LowerRange')
# 检查是否是以太网
if lowerRange[0] == 'ethernet':
driverDesc = _winreg.QueryValueEx(hSubKey, 'DriverDesc')[0]
print '检测到网络适配器名：', driverDesc
netCfgInstanceID = _winreg.QueryValueEx(hSubKey, 'NetCfgInstanceID')[0]
print '检测到网络适配器ID：', netCfgInstanceID
if netCfgInstanceID == None:
print '没有找到网络适配器，程序退出'
exit()
break
_winreg.CloseKey(hNdiInfKey)
except WindowsError:
print r'Message: No Ndi\Interfaces Key'
# 循环结束，目前只提供修改一个网卡IP的功能
_winreg.CloseKey(hSubKey)
_winreg.CloseKey(hkey)
# 通过修改注册表设置IP
strKeyName = 'System\\CurrentControlSet\\Services\\Tcpip\\Parameters\\Interfaces\\' + netCfgInstanceID
print '网络适配器的注册表地址是：\n', strKeyName
hkey = _winreg.OpenKey(_winreg.HKEY_LOCAL_MACHINE, \
strKeyName, \
0, \
_winreg.KEY_WRITE)
config = ConfigParser.ConfigParser()
print
print '正在打开IP.ini配置文件…'
config.readfp(open('IP.ini'))
IPAddress = config.get("school","IPAddress")
SubnetMask = config.get("school","SubnetMask")
GateWay = config.get("school","GateWay")
DNSServer1 = config.get("school","DNSServer1")
DNSServer2 = config.get("school","DNSServer2")
DNSServer = [DNSServer1,DNSServer2]
print '配置文件内设定的信息如下，请核对：'
print
print 'IP地址：', IPAddress
print '子关掩码：', SubnetMask
print '默认网关：', GateWay
print '主DNS服务器：', DNSServer1
print '次DNS服务器：', DNSServer2
print
res = raw_input('现在，请您决定：输入1，则将配置文件写入系统；输入2，则将现有的系统设定还原为全部自动获取；否则程序退出：')
if str(res) == '1':
try:
_winreg.SetValueEx(hkey, 'EnableDHCP', None, _winreg.REG_DWORD, 0x00000000)
_winreg.SetValueEx(hkey, 'IPAddress', None, _winreg.REG_MULTI_SZ, [IPAddress])
_winreg.SetValueEx(hkey, 'SubnetMask', None, _winreg.REG_MULTI_SZ, [SubnetMask])
_winreg.SetValueEx(hkey, 'DefaultGateway', None, _winreg.REG_MULTI_SZ, [GateWay])
_winreg.SetValueEx(hkey, 'NameServer', None, _winreg.REG_SZ, ','.join(DNSServer))
except WindowsError:
print 'Set IP Error'
exit()
_winreg.CloseKey(hkey)
print '切换成功！重置网络后即可生效'
elif str(res) == '2':
try:
_winreg.SetValueEx(hkey, 'EnableDHCP', None, _winreg.REG_DWORD, 0x00000001)
_winreg.SetValueEx(hkey, 'T1', None, _winreg.REG_DWORD, 0x00000000)
_winreg.SetValueEx(hkey, 'T2', None, _winreg.REG_DWORD, 0x00000000)
_winreg.SetValueEx(hkey, 'NameServer', None, _winreg.REG_SZ, None)
_winreg.SetValueEx(hkey, 'DhcpConnForceBroadcastFlag', None, _winreg.REG_DWORD, 0x00000000)
_winreg.SetValueEx(hkey, 'Lease', None, _winreg.REG_DWORD, 0x00000000)
_winreg.SetValueEx(hkey, 'LeaseObtainedTime', None, _winreg.REG_DWORD, 0x00000000)
_winreg.SetValueEx(hkey, 'LeaseTerminatesTime', None, _winreg.REG_DWORD, 0x00000000)
except WindowsError:
print 'Set IP Error'
exit()
_winreg.CloseKey(hkey)
print '切换成功！重置网络后即可生效'
else:
print '用户手动取消，程序退出'
exit('')
