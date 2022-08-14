# Weather Project

## Project Setup
 
 Go Version : 1.18 <br />
 FrameWork Used : Gin  

 ### Run Server 

 ```
 go run main.go 
 ```

 ### Install Dependency 

 ```
  go mod tidy 
 ```

### Project Architecture : [HexaGonal Architecture](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)) 

```bash
├── cmd
├── core
│   ├── app
│   ├── domain
│   │   └── rf_data.go
│   ├── persistence
│   │   └── rf_data.go
│   └── service
│       └── rfservice
│           └── Rfservice.go
├── gateways
├── go.mod
├── go.sum
├── handler
│   └── rfHandler
│       ├── api_handler.go
│       └── rfhandler
│           └── rfhandler.go
├── main.go
├── middleware
├── model
│   └── rf_data.go
├── readme.md
├── util
│   ├── context.go
│   ├── dbconnection.go
│   ├── env.go
│   ├── logger
│   │   └── logger.go
│   └── notification
│       └── notification.go
└── utilities
    └── constant.go
```