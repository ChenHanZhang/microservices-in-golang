package main

import (
	pb "github.com/ChenHanZhang/microservices-in-golang/proto/consignment"
	micro "github.com/micro/go-micro"

	"context"
	"fmt"
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
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	// 接收承运的货物
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	allConsignments := s.repo.GetAll()
	res.Consignments = allConsignments
	return nil
}

func main() {
	repo := Repository{}

	// 注意，这里采用 go-micro 来实现
	srv := micro.NewService(
			micro.Name("go.micro.srv.consignment"),
			micro.Version("latest"),
			)

	srv.Init()

	err := pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})

	if err = srv.Run(); err != nil {
		fmt.Println(err)
	}
}