## Go microservice using Go Kit
This is a project that takes advantage of the `go-kit` framework and creates microservices.

### Installation
#####Clone this repository in your `GOPATH` usually at 
  - `/Users/<your-username>/go/src/` for `Mac`
  - `/home/<your-username>/go/src/` for `Linux`


#####Setting up the database
- Create database `coinsph` for the user `postgres`
- Create a table `accounts`
```
CREATE TABLE accounts (Id VARCHAR (100) UNIQUE NOT NULL, username VARCHAR (50) UNIQUE NOT NULL, currency VARCHAR (20) NOT NULL, balance VARCHAR (100) NOT NULL);
```

- Create another table `payments`
```
CREATE TABLE payments (id VARCHAR(100) UNIQUE NOT NULL, sender VARCHAR(100) NOT NULL, reciever VARCHAR(100) NOT NULL, amount NUMERIC NOT NULL, initiated TIME NOT NULL, completed TIME NOT NULL);
```

#####Run server
`go run main.go`


### Code Linting
The linter used here is `golangci-lint` and to run it, simple go within the root of the repo and run `golangci-lint run`