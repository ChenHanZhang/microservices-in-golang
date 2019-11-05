module vessel-service

go 1.13

require (
	github.com/ChenHanZhang/microservices-in-golang/proto/vessel v0.0.0
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/json-iterator/go v1.1.8 // indirect
	github.com/marten-seemann/qtls v0.4.1 // indirect
	github.com/micro/go-micro v1.15.1
	github.com/miekg/dns v1.1.22 // indirect
	github.com/nats-io/nats.go v1.9.1 // indirect
	github.com/nats-io/nkeys v0.1.2 // indirect
	go.uber.org/zap v1.12.0 // indirect
	golang.org/x/crypto v0.0.0-20191029031824-8986dd9e96cf // indirect
	golang.org/x/net v0.0.0-20191101175033-0deb6923b6d9
	golang.org/x/sys v0.0.0-20191104094858-e8c54fb511f6 // indirect
	golang.org/x/tools v0.0.0-20191101200257-8dbcdeb83d3f // indirect
	google.golang.org/genproto v0.0.0-20191028173616-919d9bdd9fe6 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)

replace github.com/ChenHanZhang/microservices-in-golang/proto/vessel => ../proto/vessel

replace github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.7.1-0.20190913061013-f15a82d3fdc3
