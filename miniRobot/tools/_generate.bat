@echo OFF
chcp  65001
@echo "-----------fix package name(本地化)------------------"
py  .\amend.py
timeout 1
md ..\msg\go
@echo "-----------Proto-file(待处理)------------------"
echo _generate.bat path : %~dp0
dir    %~dp0\..\msg\proto\*.proto /B > list.txt              
REM '待处理的Proto文件'
for  /f  %%a  in  (list.txt)  do (
echo 正在转换 %%a  
protoc -I=%~dp0\..\msg\proto\ --go_out=..\msg\go %%a
echo 忙碌中...
)

@echo "------------Go-file(已生成)--------------------"
for /R "..\msg\go" %%s in (*.go) do (@echo "creating->file:%%s")

@echo "------------c++代码(协议注册)--------------------"
py  .\convertCpp.py

@echo "------------若无操作 3秒后自动退出--------------------"
timeout 3
Exit