``` sh
Clone with ssh recommended
$ git clone git@github.com:AndresXLP/backend-tweets.git

Clone with https
$ git clone https://github.com/AndresXLP/backend-tweets.git 
```

# Requirements

* go v1.18
* go modules

# Technology Stack

- Framework: [echo](https://echo.labstack.com/)
- Validations: [validator](https://github.com/go-playground/validator)
- Encrypt Password: [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- JWT: [Jwt-Go](https://github.com/golang-jwt/jwt)

# Architecture

- [Hexagonal Architecture](https://www.happycoders.eu/software-craftsmanship/hexagonal-architecture/)

# Build

* Install dependencies:

```sh
$ go mod download
```

* Run with Docker:

```sh 
$ make compose-up 
```

# Environments

#### Required environment variables

* `SERVER_HOST`: host for the server
* `SERVER_PORT`: port for the server
* `DB_HOST`: host database
* `DB_PORT`: port database
* `DB_USER`: user database
* `DB_PASSWORD`: password database
* `DB_NAME`: name database
* `SECRET_JET`: secret for JWT


# Contributors

* Andres Puello

