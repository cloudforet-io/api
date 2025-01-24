module github.com/cloudforet-io/api/dist

go 1.23.4

replace github.com/cloudforet-io/api => ./

require (
	github.com/golang/protobuf v1.5.4
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.26.0
	google.golang.org/genproto/googleapis/api v0.0.0-20250122153221-138b5a5a4fd4
	google.golang.org/grpc v1.70.0
	google.golang.org/protobuf v1.36.3
)

require (
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
)
