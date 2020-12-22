@echo off
set copy__path=E:\go_prj\go_project\miniRobot
set target_path=C:\Users\123\AppData\Local\Packages\CanonicalGroupLimited.UbuntuonWindows_79rhkp1fndgsc\LocalState\rootfs\home\pitter\goprj\miniRobot\
xcopy %copy__path%*.* %target_path% /s /h /y /U
REM 仅拷贝更新目标目录里存在的文件 /U
REM 旧版本
REM xcopy %copy__path%*.* %target_path% /s /h /y /e

REM 以下是新版本
REM  需要过滤的文件或目录填至EXCLUDE.txt文件内 每个文件或者文件夹占一行
REM 01-30-2020之前的文件都被过滤掉 %copy__path%目录后不能由斜杠
cd /d %~dp0
.\xcopydate.bat %copy__path% %target_path% 01-30-2020
Exit
