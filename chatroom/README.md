# Architecture
![image](https://github.com/Glassylip/PracticeItems/blob/main/chatroom/architecture.png)

# Quick Start
## Start Redis(Need to install redis)
```bash
redis-server
```

## Start Server
```bash
cd chatroom\server\main
go run main.go processor.go redis.go
```

## Start Client
```bash
ch chatroom\client\main
go run main.go
```
