[![Build Status](https://github.com/tischda/with-time/actions/workflows/build.yml/badge.svg)](https://github.com/tischda/with-time/actions/workflows/build.yml)
[![Test Status](https://github.com/tischda/with-time/actions/workflows/test.yml/badge.svg)](https://github.com/tischda/with-time/actions/workflows/test.yml)
[![Coverage Status](https://coveralls.io/repos/tischda/with-time/badge.svg)](https://coveralls.io/r/tischda/with-time)
[![Go Report Card](https://goreportcard.com/badge/github.com/tischda/with-time/v3)](https://goreportcard.com/report/github.com/tischda/with-time/v3)

# with-time


Windows utility that injects a %TIME:format% environment variable in a command.

The format is a golang [Time.Format](https://pkg.go.dev/time#Time.Format).


### Install

~~~
go install github.com/tischda/with-time
~~~

### Usage

~~~
Usage: with-time "COMMAND"

Injects a %TIME:format% environment variable in a command.
COMMAND must be quoted and contain a %TIME:format% substring.

The format can be a golang Time.Format layout (e.g. 20060102) or a
YYYYMMDD style format. If format is empty, time.UnixDate is used.
(e.g. "Mon Jan _2 15:04:05 MST 2006")

  YYYY -> 2006 (Year)
  MM   -> 01   (Month)
  DD   -> 02   (Day)
  HH   -> 15   (Hour)
  mm   -> 04   (Minute)
  ss   -> 05   (Second)

OPTIONS:

  -?, --help
          display this help message
  -v, --version
          print version and exit
~~~

## Examples

~~~
$ with-time.exe "echo now is %TIME:20060102-150405% and windir is %windir%"
now is 20170720-150748 and windir is C:\WINDOWS
~~~

Note that in a batch script the `%` surrounding the TIME variable must be doubled:

~~~
@echo off
with-time.exe "echo now is %%TIME:20060102-150405%% and windir is %windir%"
~~~

Assign the TIME value to another variable (only works in a batch script):
~~~
@echo off
for /f "delims=" %%a in ('with-time "echo %%TIME:20060102-150405%%"') do set NOW=%%a
echo %NOW%
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
