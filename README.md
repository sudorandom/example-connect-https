# Connect w/HTTPS Example
This is a simple Go example that shows how to run a ConnectRPC server with TLS (not relying on h2c) and connect to it with a client.

## Running the server
```shell
go run server/main.go
```

## Running the local client
```shell
go run client/main.go
```

## Running the client against [demo.connectrpc.com](https://demo.connectrpc.com/)
```shell
go run demo-client/main.go
```

## Certificates
The local certificates were created with this command:
```shell
openssl req -new -newkey rsa:4096 -days 365 -nodes -x509 \
                                                   -subj "/C=DK/L=Copenhagen/O=kmcd/CN=local.kmcd.dev" \
                                                   -keyout cert.key  -out cert.crt
```