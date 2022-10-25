# cli-template

<p align="center">
  <img width="120" src="./cli-template.svg">
</p>

<a href="https://opensource.org/licenses/MIT">
    <img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License: MIT">
</a>
<a href="https://meson.network/">
    <img src="https://img.shields.io/badge/made%20by-Meson%20Network-blue.svg?style=flat-square" alt="License: MIT">
</a>

> This is a template project that you can fork or copy as a project to start with.

## ⚡️ Designed With

- Easy Configuration (toml)
- Command Line Support
- Rich Plugins

## 🎈 Main Entry Point File

```bash
main.go is the entry file
other test entries are under `test` folders
```

## 📚 Steps To Start

### Step1. Config Your Project

- Check folder `root_conf` where you can put different {config_name}.toml files
- The 'default.toml' config is used if not explicited configured 
- You can run `go run ./ config set --https.enable=false` to setup your own config file (generated inside 'user_conf/default.toml' )
- You can run 'go run ./ config --conf=test set --https.enable=false' to setup your own config file (generated inside 'user_conf/test.toml' )
- You can also edit user_conf/*.toml files directly without the help of command line


### Step2. Check Database Initialization

- Config your database in your *.toml file
- Open your database and construct the tables using file `cmd_db/table.sql`
- Run `go run ./ db init` which will call the function `Initialize()` inside the `cmd_db/initialize.go` file which initialize the db data

### Step3. Write API

- Go to `cmd_default` folder where your main program locate
- Go to `http/api` folder to add your own api file
- Run `go run ./ gen_api` to auto generate your api files
- Run `go run ./` to start your main program with http server 
- You can view the api pages with `localhost`

## ⚙️ Command line hints

```bash
$ go run . gen_api                //generate api
$ go run . config set ...         //set configs
$ go run . log                    //show all logs
$ go run . log --only_err=true    //show all error logs [error,panic,fatal]
```