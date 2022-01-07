# CliTemplate


# How to use
```
go build 

1."default_" application:
default_ is the main program

1.1 run default program with dev mode
./CliAppTemplate --dev=true

1.2 run default program with production mode
./CliAppTemplate  
or 
./CliAppTemplate --dev=false


2."config" application:
config is the program used to show or set config file

2.1 set dev.json config
./CliAppTemplate --dev=true set ... 

2.1 set pro.json config
./CliAppTemplate --dev=false set ... 

3.run log application 
log is used to show the local log files

3.1 show all logs
./CliAppTemplate log

3.2 only show error logs : [error,panic,fatal]
./CliAppTemplate log --onlyerr=true
 
4."service" application:
service is used to set application to OS service 


```


## Running process
```sh

1.entry -> main.go
2.basic logger is initialized 
3.cmd/cmd.go ->ConfigCmd() is called
4.check "dev" mode or "pro" mode 
5.read the related "dev.json" or "pro.json" config file
6.--> go to cmd application "config"|"default_"|"log"|"service"



```