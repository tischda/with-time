# gonow [![Build status](https://ci.appveyor.com/api/projects/status/nb8spqtr99has88f?svg=true)](https://ci.appveyor.com/project/tischda/gonow)

Windows utility written in [Go](https://www.golang.org) to inject a %NOW% environment variable in a command.

### Install

There are no dependencies.

~~~
go get github.com/tischda/gonow
~~~

### Usage

~~~
Usage: gonow [-h] [-v] command

OPTIONS:
  -h
  -help
        displays this help message
  -v
  -version
        print version and exit
~~~

## Examples

~~~
$ gonow.exe "echo now is %NOW% and windir is %windir%"
now is 20170720-150748 and windir is C:\WINDOWS
~~~

Note that in a batch script the `%` must be doubled:

~~~
@echo off
gonow.exe "echo now is %%NOW%% and windir is %windir%"
~~~

Assign the NOW value to another variable:
~~~
$ for /f "delims=" %a in ('gonow "echo %NOW%"') do @set REALNOW=%a
$ echo %REALNOW%
20170720-151620
~~~

and in a batch script:
~~~
@echo off
for /f "delims=" %%a in ('gonow "echo %%NOW%%"') do set REALNOW=%%a
echo %REALNOW%
~~~


### Background

I used to write this kind of code to name my log files with the current date:

~~~
:: compute current date for log file
set T=%TIME: =0%
set T=%T:~0,2%%T:~3,2%%T:~6,2%
set D=%DATE:~-4%%DATE:~-7,-5%%DATE:~-10,-8%
set LOGFILE=%LOG_DIR%\%JOB%_%D%-%T%.log
~~~

But this method only works for the French locale and is generally error prone.


### References

* https://stackoverflow.com/questions/1192476/format-date-and-time-in-a-windows-batch-script
