# URL shortener using go

## First install redis 
Installation guide -> [link](https://redis.io/docs/install/install-redis/)

## Install necessary go modules
```
rogerding@CPU url-shortener-go % pwd
$HOME/VSCode/GIT/url-shortener-go
rogerding@CPU url-shortener-go % go mod tidy
```

## Start up redis server on terminal
1. Open terminal
```
rogerding@CPU ~ % redis-server
17567:C 13 Nov 2023 10:51:00.934 * oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
17567:C 13 Nov 2023 10:51:00.934 * Redis version=7.2.3, bits=64, commit=00000000, modified=0, pid=17567, just started
17567:C 13 Nov 2023 10:51:00.934 # Warning: no config file specified, using the default config. In order to specify a config file use redis-server /path/to/redis.conf
17567:M 13 Nov 2023 10:51:00.935 * Increased maximum number of open files to 10032 (it was originally set to 256).
17567:M 13 Nov 2023 10:51:00.935 * monotonic clock: POSIX clock_gettime
                _._
           _.-``__ ''-._
      _.-``    `.  `_.  ''-._           Redis 7.2.3 (00000000/0) 64 bit
  .-`` .-```.  ```\/    _.,_ ''-._
 (    '      ,       .-`  | `,    )     Running in standalone mode
 |`-._`-...-` __...-.``-._|'` _.-'|     Port: 6379
 |    `-._   `._    /     _.-'    |     PID: 17567
  `-._    `-._  `-./  _.-'    _.-'
 |`-._`-._    `-.__.-'    _.-'_.-'|
 |    `-._`-._        _.-'_.-'    |           https://redis.io
  `-._    `-._`-.__.-'_.-'    _.-'
 |`-._`-._    `-.__.-'    _.-'_.-'|
 |    `-._`-._        _.-'_.-'    |
  `-._    `-._`-.__.-'_.-'    _.-'
      `-._    `-.__.-'    _.-'
          `-._        _.-'
              `-.__.-'

17567:M 13 Nov 2023 10:51:00.937 # WARNING: The TCP backlog setting of 511 cannot be enforced because kern.ipc.somaxconn is set to the lower value of 128.
17567:M 13 Nov 2023 10:51:00.938 * Server initialized
17567:M 13 Nov 2023 10:51:00.938 * Loading RDB produced by version 7.2.3
17567:M 13 Nov 2023 10:51:00.938 * RDB age 157659 seconds
17567:M 13 Nov 2023 10:51:00.938 * RDB memory usage when created 1.21 Mb
17567:M 13 Nov 2023 10:51:00.938 * Done loading RDB, keys loaded: 1, keys expired: 2.
17567:M 13 Nov 2023 10:51:00.938 * DB loaded from disk: 0.001 seconds
17567:M 13 Nov 2023 10:51:00.938 * Ready to accept connections tcp
```
2. Open another tab in terminal and confirm redis is accessible
```
rogerding@CPU ~ % redis-cli
127.0.0.1:6379>
```

## Run main go application
1. In project directory start up the main go script
```
rogerding@CPU url-shortener-go % go run main.go 
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> main.main.func1 (3 handlers)
[GIN-debug] POST   /create-shortened-url     --> main.main.func2 (3 handlers)
[GIN-debug] GET    /:shortenedUrl            --> main.main.func3 (3 handlers)

Redis client started successfully
Pong message = {PONG}[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
```
2. In new tab on terminal test out url shortener
```
rogerding@CPU ~ % curl -X POST --data '{"source_url": "https://www.google.com"}' http://localhost:8080/create-shortened-url
{"message":"Shortened url created successfully!","shortened_url":"http://localhost:8080/rGu2aeQORK"}%
```
3. Verify in redis that entry is stored
```
rogerding@CPU ~ % redis-cli
127.0.0.1:6379> KEYS *
1) "rGu2aeQORK"
```
4. Verify that the shortened URL redirects to the original source url

# That's it! Simple and straightforward :)
