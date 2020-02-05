@echo OFF
REM '待处理的Proto文件'
@echo "-----------Proto-file------------------"
set protocpath=C:\Users\Pitter\go\bin
for /R ".\" %%i in (*.proto) do (%protocpath%\protoc -I=%~dp0 --go_out=..\go  %%i
@echo "exec->file:%%i ")
REM '生成的golang文件'
@echo "------------Go-file--------------------"
for /R "..\go" %%s in (*.go) do (@echo "creating->file:%%s")

REM '区分文件夹和文件'
REM   timeout 3
    REM for /f "delims=" %%i in ('dir /a/b/s') do pushd "%%i" 2>nul && (call :folder "%%i" & popd) || call :file "%%i"
    REM pause
REM goto :eof
REM :file
	REM echo creating: %~1 file is OK!
REM goto :eof
REM :folder
	REM REM echo %~1 is Folder!
REM goto :eof
REM PAUSE
timeout 3
Exit