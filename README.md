# Example gRPC server on Go on Server Streaming

## What is this?
This is a sample client-server gRPC app where the server would return a random number every n seconds, n is between 0 and 10.

## Usage

Start the server:
```
go run cmd/server/main.go
```

Start the client:
```
go run cmd/client/main.go
```