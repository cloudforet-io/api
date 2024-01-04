module github.com/cloudforet-io/api/dist

go 1.21.4

replace github.com/cloudforet-io/api => ./

require (
	github.com/golang/protobuf v1.5.3
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.0
	google.golang.org/genproto/googleapis/api v0.0.0-20240102182953-50ed04b92917
	google.golang.org/grpc v1.60.1
	google.golang.org/protobuf v1.32.0
)

require (
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20231212172506-995d672761c0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240102182953-50ed04b92917 // indirect
)
