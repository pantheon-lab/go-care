module greeter

go 1.19

require (
	github.com/pantheon-lab/go-care v1.1.0
	google.golang.org/grpc v1.50.1
	proto v0.0.0
)

replace proto v0.0.0 => ./proto

replace github.com/pantheon-lab/go-care v1.1.0 => ../../

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
