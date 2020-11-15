@echo off
set copy__path=E:\go_prj\go_project\miniRobot\
set target_path=C:\Users\123\AppData\Local\Packages\CanonicalGroupLimited.UbuntuonWindows_79rhkp1fndgsc\LocalState\rootfs\home\pitter\goprj\miniRobot\
xcopy %copy__path%*.* %target_path% /s /h /y /U
REM 仅拷贝更新目标目录里存在的文件 /U
timeout 3
Exit