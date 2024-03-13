module github.com/cloudforet-io/api/dist

go 1.22.1

replace github.com/cloudforet-io/api => ./

require (
	github.com/golang/protobuf v1.5.4
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.1
	google.golang.org/genproto/googleapis/api v0.0.0-20240311173647-c811ad7063a7
	google.golang.org/grpc v1.62.1
	google.golang.org/protobuf v1.33.0
)

require (
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240311132316-a219d84964c2 // indirect
)
