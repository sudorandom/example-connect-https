module github.com/sudorandom/example-connect-https

go 1.22.4

require (
	buf.build/gen/go/connectrpc/eliza/connectrpc/go v1.16.2-20230913231627-233fca715f49.1
	buf.build/gen/go/connectrpc/eliza/protocolbuffers/go v1.34.2-20230913231627-233fca715f49.2
	connectrpc.com/connect v1.16.2
)

require google.golang.org/protobuf v1.34.2 // indirect
