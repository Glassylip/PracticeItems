フォルダー パスの一覧
ボリューム シリアル番号は 0C7D-0E6A です
C:.
│  go.mod
│  go.sum
│  README.md
│  t.txt
│  
├─client
│  │  login.go
│  │  main.go
│  │  utils.go
│  │  
│  ├─main
│  │      main.go
│  │      
│  ├─model
│  │      curUser.go
│  │      
│  ├─process
│  │      server.go
│  │      smsMgr.go
│  │      smsProcess.go
│  │      userMgr.go
│  │      userProcess.go
│  │      
│  └─utils
│          utils.go
│          
├─common
│  └─message
│          message.go
│          user.go
│          
└─server
    ├─main
    │      main.go
    │      processor.go
    │      redis.go
    │      
    ├─model
    │      error.go
    │      user.go
    │      userDao.go
    │      
    ├─process
    │      smsProcess.go
    │      userMgr.go
    │      userProcess.go
    │      
    ├─service
    └─utils
            utils.go
            
