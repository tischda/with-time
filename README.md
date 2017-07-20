# go-now [![Build status](https://ci.appveyor.com/api/projects/status/apwc7sg9sak0syjx?svg=true)](https://ci.appveyor.com/project/tischda/go-now)

Windows utility written in [Go](https://www.golang.org) to inject a %NOW% environment variable in a command.

### Install

There are no dependencies.

~~~
go get github.com/tischda/go-now
~~~

### Usage

~~~
Usage: go-now [-h] [-v] command

OPTIONS:
  -h
  -help
        displays this help message
  -v
  -version
        print version and exit
~~~

Example:

~~~
# go-now "echo now is %%NOW%% and windir is %windir%"
now is 20170720-150748 and windir is C:\WINDOWS
~~~

Note that the `%%NOW%%` is replaced internally so `%` must be doubled.

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

OK, but I still want to set a variable. Now you can do this:

~~~
# for /f "delims=" %%a in ('go-now "echo %%%%NOW%%%%"') do @set REALNOW=%%a
# echo %REALNOW%
20170720-151620
~~~

### References

* https://stackoverflow.com/questions/1192476/format-date-and-time-in-a-windows-batch-script
