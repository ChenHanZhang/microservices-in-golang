package main

import (
	"fmt"
	micro "github.com/micro/go-micro"
	"log"
	"os"

	pb "github.com/ChenHanZhang/microservices-in-golang-proto/consignment"
	vesselProto "github.com/ChenHanZhang/microservices-in-golang-proto/vessel"
)

const (
	defaultHost = "localhost:27017"
)

func main() {
	// Database host from the environment variables
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}
	session, err := CreateSession(host)

	if err != nil {
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	// 确保在main退出前关闭会话
	defer session.Close()

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)
	vesselClient := vesselProto.NewVesselServiceClient("go.micro.srv.vessel", srv.Client())
	// Init will parse the command line flags.
	srv.Init()
	// Register handler
	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})
	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}