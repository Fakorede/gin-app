# gin-app
An app built with Golang and the Gin Framework

### Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Installation
$ git clone https://github.com/Fakorede/gin-app
$ cd gin-app
$ go get
$ go run server.go

### Heroku Deployment

Create a `Procfile` and add below command,

```
web: bin/golang-gin-poc
```

Run command:

```
go build -o bin/golang-gin-poc -v .
```
