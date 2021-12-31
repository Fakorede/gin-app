# gin-app
An app built with Golang and the Gin Framework

### Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Installation

```
$ git clone https://github.com/Fakorede/gin-app
$ cd gin-app
$ go mod download
```

### Setup env variables
Create a `.env` file in the root of the project and add the ff variables:

```
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=gin_app
DB_USER=
DB_PASS=

JWT_SECRET=
JWT_ISSUER=

AUTH_USERNAME=
AUTH_PASSWORD=

PORT=5000
```

### Run App

```
$ go run server.go
```

### Docs

The api documentation is available at `http://localhost:5000/swagger/index.html`
