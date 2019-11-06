module consignment-service

go 1.13

require (
	github.com/micro/go-micro v1.15.1
	github.com/miekg/dns v1.1.22 // indirect
	golang.org/x/crypto v0.0.0-20191029031824-8986dd9e96cf // indirect
	golang.org/x/net v0.0.0-20191101175033-0deb6923b6d9
	golang.org/x/sys v0.0.0-20191104094858-e8c54fb511f6 // indirect
)

require (
	github.com/ChenHanZhang/microservices-in-golang-proto v1.0.1
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)

replace github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.7.1-0.20190913061013-f15a82d3fdc3
