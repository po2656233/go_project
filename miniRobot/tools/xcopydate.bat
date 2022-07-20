@echo off
rem transfers all files in all subdirectories of
rem the source drive or directory (%1) to the destination
rem drive or directory (%2)
rem  lastModifyTime (%3)

xcopy %1 %2 /s /d:%3 /exclude:.\EXCLUDE.txt /y
if errorlevel 4 goto lowmemory
if errorlevel 2 goto abort
if errorlevel 0 goto exit
:lowmemory
echo Insufficient memory to copy files or
echo invalid drive or command-line syntax.
goto exit
:abort
echo You pressed CTRL+C to end the copy operation.
goto exit
:exit
timeout 3