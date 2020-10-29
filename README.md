# golang-grpc-example

### generate gRPC code

```
$ protoc -I ./proto pinger.proto --go_out=plugins=grpc:./pinger
```

### Usage

```
$ go run server/main.go
```

```
$ go run client/main.go
```
