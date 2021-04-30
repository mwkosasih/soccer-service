## soccer-service

### Authors 
- Mohammad Wildan Kosasih <mwkosasih@gmail.com>

### About
Service for manage data soccer (teams & players). 

### How
How to run this project :
- clone this project into go path 
- run in terminal `go run main.go`

### Structure
```
├── action
├── builder
├── entity
├── repo
    ├── mysql
    ├── redis
├── transport
    ├── grpc
├── util             
├── main.go
```
- Action: package for orchestrate and perform domain actions
- Builder: package for handling data transform
- Entity: internal entity struct
- Repo: package for handling business and data process
- Util: common app process / helper
- Transport: Server connection request, package for handling request. 

### Testing  
How to run unit test :
- go test -cover -coverprofile=coverage.out ./...

check coverage result :
- go tool cover -html=coverage.out
