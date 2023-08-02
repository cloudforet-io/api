module github.com/cloudforet-io/api/dist

go 1.20

replace github.com/cloudforet-io/api => ./

require (
	github.com/golang/protobuf v1.5.3
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.2
	google.golang.org/genproto/googleapis/api v0.0.0-20230731193218-e0aa005b6bdf
	google.golang.org/grpc v1.57.0
	google.golang.org/protobuf v1.31.0
)

require (
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	google.golang.org/genproto v0.0.0-20230726155614-23370e0ffb3e // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230731190214-cbb8c96f2d6d // indirect
)
