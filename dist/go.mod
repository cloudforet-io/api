module github.com/cloudforet-io/api/dist

go 1.22.4

replace github.com/cloudforet-io/api => ./

require (
	github.com/golang/protobuf v1.5.4
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.20.0
	google.golang.org/genproto/googleapis/api v0.0.0-20240701130421-f6361c86f094
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.34.2
)

require (
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240624140628-dc46fd24d27d // indirect
)
