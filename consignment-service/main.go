package main

import (
	pb "github.com/ChenHanZhang/microservices-in-golang/proto/consignment"
	vesselProto "github.com/ChenHanZhang/microservices-in-golang/proto/vessel"
	micro "github.com/micro/go-micro"
	"log"

	"context"
	"fmt"
)

const (
	PORT = ":50051"
)

type Repository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error) // 存放新货物
	GetAll() []*pb.Consignment
}

// 存放多批货物的仓库，实现 IRepository 接口
type ConsignmentRepository struct {
	consignments []*pb.Consignment
}


// 相当于定义类方法
func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

func (repo *ConsignmentRepository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// 定义微服务, 要实现 ShippingServer 接口
type service struct {
	// 这里使用接口，TODO: 其实直接用 ConsignmentRepository 也是可以的
	repo Repository

	// 这里注册(记录)一个货船服务的客户端对象
	vesselClient vesselProto.VesselServiceClient
}

// service 实现 consignment.pb.go 中的 ShippingServiceServer 接口
// 使 service 作为 gRPC 的服务端
// 就是这个CreateConsignment方法，它接受一个context以及proto中定义的
// TODO: Consignment消息，这个Consignment是由 gRPC 的服务器处理后提供给你的
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	// 这里，我们通过货船服务的客户端对象 TODO: 客户端对象?和之前在cli中的有什么不同
	// 向货船服务发起一个请求
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		Capacity:             req.Weight,
		MaxWeight:            int32(len(req.Containers)),
	})

	if err != nil {
		return err
	}

	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)

	req.VesselId = vesselResponse.Vessel.Id

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
	repo := &ConsignmentRepository{}

	// 注意，这里采用 go-micro 来实现
	srv := micro.NewService(
			micro.Name("consignment"),
			micro.Version("latest"),
			)

	// 在这里使用预置的方法生成一个货船服务的客户端对象
	vesselclient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())

	srv.Init()

	// TODO: 这里的 service 和 pb 的(依赖)关系
	err := pb.RegisterShippingServiceHandler(srv.Server(), &service{repo, vesselclient})

	if err = srv.Run(); err != nil {
		fmt.Println(err)
	}
}