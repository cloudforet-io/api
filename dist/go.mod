module github.com/cloudforet-io/api/dist

go 1.24.3

replace github.com/cloudforet-io/api => ./

require (
	github.com/golang/protobuf v1.5.4
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.26.3
	google.golang.org/genproto/googleapis/api v0.0.0-20250519155744-55703ea1f237
	google.golang.org/grpc v1.72.2
	google.golang.org/protobuf v1.36.6
)

require (
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250512202823-5a2f75b736a9 // indirect
)
