echo off & color 0A

set DIR="%cd%"
set PROTOC=protoc.exe
echo DIR=%DIR%
REM set GoPath=E:\temp_prj\go
REM echo DIR=%DIR% GoPath=%GoPath%
echo                START
echo -------------proto-file------------------
for /R %DIR% %%f in (*.proto) do ( 
    echo %%f 
    %PROTOC% --go_out=%DIR% -I=%DIR% %%f
	REM %PROTOC% --go_out=%GoPath% -I=%DIR% %%f REM 生成至指定目录
	rem %PROTOC% -I=%DIR% %%f --go_out=plugins=grpc:.
)

echo ------------create:go-file--------------------
for /R %DIR% %%g in (*.go) do ( 
    echo %%g
)

echo "--------------END-------------------"
REM timeout 3
pause
Exit