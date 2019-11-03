package main

import (
	pb "github.com/ChenHanZhang/microservices-in-golang/proto/consignment"

	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	PORT = ":50051"
)

type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error) // 存放新货物
}

// 存放多批货物的仓库，实现 IRepository 接口
type Repository struct {
	consignments []*pb.Consignment
}


// 相当于定义类方法
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// 定义微服务, 要实现 ShippingServer 接口
type service struct {
	repo Repository
}

// service 实现 consignment.pb.go 中的 ShippingServiceServer 接口
// 使 service 作为 gRPC 的服务端
// 就是这个CreateConsignment方法，它接受一个context以及proto中定义的
// TODO: Consignment消息，这个Consignment是由 gRPC 的服务器处理后提供给你的
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	// 接收承运的货物
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	resp := &pb.Response{
		Created:              true,
		Consignment:          consignment,
	}
	return resp, nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	allConsignments := s.repo.GetAll()
	resp := &pb.Response{Consignments: allConsignments}
	return resp, nil
}

func main() {

	utils.PrintText("running...")

	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listen on: %s\n", PORT)

	server := grpc.NewServer()
	repo := Repository{}

	// 向 rRPC 服务器注册微服务
	// 此时会把我们自己实现的微服务 service 与协议中的 ShippingServiceServer 绑定
	pb.RegisterShippingServiceServer(server, &service{repo})
	// 上述函数的原型 RegisterShippingServiceServer(s *grpc.Server, srv ShippingServiceServer)
	// 其实就是将 一个实现了 ShippingServiceServer 的实例绑定到 server 上
	// TODO: 具体的内部实现细节暂且不考虑

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}