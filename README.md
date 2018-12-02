# hashing-server

## How to run

Run the tests with

```
go test
```
To run the server, build a binary and execute it.

```
go build
./hashing-server
```

By default the server listens on port `8080`. You can override this by specifying a `PORT` environment variable.

```
PORT=4321 ./hashing-server
```

## Assumptions

* We're using Go version 1.8 or higher. I take advantage of the ![`Server.Shutdown()`](https://golang.org/pkg/net/http/#Server.Shutdown) functionality introduced in that version to perform the graceful shutdown.
* I implemented the requirement to only reply after 5 seconds for non-error cases only. If an invalid request is sent (wrong HTTP method, body missing `password` param, etc.), the server responds with an error immediately.

## Production considerations

Before deploying this to a production environment, I would consider the following changes.

* Ensure this server can use SSL.
* Use a logging library which allows different output formats, such as JSON, allowing logs to be machine readable for a log aggregation service.
* Use a robust configuration library which allows command line flags and/or environment variables, default values, type validation, etc.
