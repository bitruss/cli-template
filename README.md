# CliTemplate


# How to use
```
go build 

1."default_" application:
default_ is the main program

1.1 run default program with dev mode
./CliAppTemplate --mode=dev or  go run . --mode=dev 

1.2 run default program with production mode
./CliAppTemplate or  go run . --mode=pro 

1.3 if you want to see the config details
 go run . --conf=show


2."config" application:
config is the program used to show or set config file

2.1 set dev.json config
./CliAppTemplate --mode=dev set ... 

2.1 set pro.json config
./CliAppTemplate --mode=pro set ... 
or 
./CliAppTemplate set ...

3.run log application 
log is used to show the local log files

3.1 show all logs
./CliAppTemplate log

3.2 only show error logs : [error,panic,fatal]
./CliAppTemplate log --only_err=true
 
4."service" application:
service is used to set application to OS service 

5. "api" application:
5.1 generate the api documents
run ./ gen_api 


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