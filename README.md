# hashing-server

## How to run

Run the tests with

```
go test
```

Start the server with

```
go run main.go
```

Or you can build a binary and execute it.

```
go build
./hashing-server
```

By default the server listens on port `8080`. You can override this by specifying a `PORT` environment variable.

```
PORT=4321 go run main.go
```

## Production considerations

Before deploying this to a production environment, I would consider the following changes.

* Ensure this server can use SSL.
* Use a logging library which allows different output formats, such as JSON, allowing logs to be machine readable for a log aggregation service.
* Use a robust configuration library which allows command line flags and/or environment variables, default values, type validation, etc.
