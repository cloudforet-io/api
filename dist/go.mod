module github.com/cloudforet-io/api/dist

go 1.20

replace github.com/cloudforet-io/api => ./

require (
	github.com/golang/protobuf v1.5.3
	google.golang.org/genproto/googleapis/api v0.0.0-20230530153820-e85fd2cbaebc
	google.golang.org/grpc v1.56.1
	google.golang.org/protobuf v1.31.0
)

require (
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230525234025-438c736192d0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230525234030-28d5490b6b19 // indirect
)
