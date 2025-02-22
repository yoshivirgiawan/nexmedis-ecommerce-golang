This system provides the basic structure for building Golang applications by:

```
* REST API
* JWT authentication
* Middleware
* Logging
* SMTP Email
```

# Structure

Project structure:

```
app
├── http
│   ├── controllers
│       ├── v1
│           ├── user_controller.go
│   ├── middlewares
│       ├── auth_middleware.go
│       ├── request_logger.go
├── modules
│   ├── auth
│       ├── service.go
│   ├── jwtgenerator
│       ├── jwtgenerator.go
│   ├── user
│       ├── entity.go
│       ├── formatter.go
│       ├── input.go
│       ├── repository.go
│       ├── service.go
cmd
├── migrate.go
├── root.go
├── seed.go
config
├── database.go
├── redis.go
helper
├── api_response.go
├── errors_log.go
├── env.go
routes
├── routes.go
├── v1
    ├── routes_v1.go
storage
├── public
├── logs
    ├── api
    ├── errors
.env.example
```

# Steps

```
1. Clone this boilerplate to your local directory.
2. Create a `.env` file in the project root directory and fill it with the database configuration and JWT.
3. Run `go mod download` to install dependencies.
4. Run `go run main.go migrate` to migrate database
5. Run `go run main.go seed` to seeding database
6. Run `go run main.go` to run the application.
```

# References

```
Golang: https://go.dev/
JWT: https://jwt.io/
```
