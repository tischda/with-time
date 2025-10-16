@echo off
setlocal 

set SCRIPT_DIR=%~dp0
pushd "%SCRIPT_DIR%"

set SUT=%SCRIPT_DIR%\..\with-time.exe

echo Expected:
echo now is %DATE%-%TIME:~0,8% and windir is C:\WINDOWS
echo.
echo Actual:
%SUT% "echo now is %%TIME:02/01/2006-15:04:05%% and windir is %windir%"

echo.

echo Expected:
echo %DATE%-%TIME:~0,8%
echo.
echo Actual:
for /f "delims=" %%a in ('%SUT% "echo %%TIME:02/01/2006-15:04:05%%"') do @set NOW=%%a
echo %NOW%

endlocal
