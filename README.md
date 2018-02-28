# go-gin-mgo-restapi-starter-project
A RESTFul API starter project by using Golang Gin Govendor and mgo
### Config by using .env
```
# create .env file to root directory (Follow from .env.example)
```
### Run it
```
# Install a govendor by using "go get"
$ go get -u github.com/kardianos/govendor
# download dependencies by using govendor
$ govendor sync
$ go run ./src/main.go ./src/routers.go
```
### Build it
```
$ govendor sync
$ go build -o ./runablefile ./src/
$ ./runablefile
```