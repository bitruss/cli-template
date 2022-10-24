# cli-template

# How to use


#### before you start first check your config

#### root_conf 
```
root_conf is your default config where you could put 
various kinds of toml files . e.g : "default.toml" , "dev.toml" ....
```


## Running process
```
.check root_conf folder 
.check db folder (you may need to deploy the .sql)
```

```
1.entry -> main.go
2.basic logger is initialized 
3.cmd/cmd.go ->ConfigCmd() is called
4.read the related config file
5.--> go to cmd application "config"|"default_"|"log"|"db"

```


#### 1."default_" sub-program:

##### default_ is the main program

#### 1.1 run default program with {config_name}

##### ```go run . --conf={config_name}``` // will use the {config_name}.toml inside configs folder

##### ```go run .```  // just use defalut.toml

#### 2."config" sub-program::

##### config is the program used to show or set config file

#### 2.1 set config

##### ```go run . --conf={config_name} config set ...```

##### ```go run . config set ...```   //using default.toml

#### 3. log sub-program:

#### 3.1 show all logs

##### ```go run . log```

#### 3.2 only show error logs : [error,panic,fatal]

##### ```go run . log --only_err=true```

#### 4. "api" sub-program::

##### 4.1 generate the api documents

##### ```go run . gen_api```


#### 5. "db" sub-program::

##### ```go run . db init_data```


## API

```
After running default_ main program, you can go directly to your browser to 
check and invoke the api swagger docs
```
