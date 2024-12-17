module github.com/cloudforet-io/api/dist

go 1.23.4

replace github.com/cloudforet-io/api => ./

require (
	github.com/golang/protobuf v1.5.4
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.24.0
	google.golang.org/genproto/googleapis/api v0.0.0-20241216192217-9240e9c98484
	google.golang.org/grpc v1.69.0
	google.golang.org/protobuf v1.36.0
)

require (
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.20.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241209162323-e6fa225c2576 // indirect
)
