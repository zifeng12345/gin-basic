# Go Gin Example 

## Installation
```
$git clone https://github.com/zifeng12345/gin-basic.git
```

## How to run

### Required

- Mysql

### Ready

run script/waiting.sql in database table

### Conf

You should modify `conf/config.toml`

```
[mysql]
host = "localhost"
port = 3308
user = "root"
passwd = "root"
db = "gin"
timeout = "10s"
...
```

### Run
```
$ cd gin-basic
$ go mod tidy
$ go run main.go 
```

Project information and existing API

```
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /api/v1/waiting

Listening port is 8000
```

## Features

- RESTful API
- Gorm
- Gin
- JWT