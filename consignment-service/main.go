package main

import (
	"errors"
	"fmt"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"golang.org/x/net/context"
	"log"
	"os"

	pb "github.com/ChenHanZhang/microservices-in-golang-proto/consignment"
	userService "github.com/ChenHanZhang/microservices-in-golang-proto/user"
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
		// 用户认证的中间件
		micro.WrapHandler(AuthWrapper),
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

// AuthWrapper 是一个高阶函数，它接受一个函数A，且返回一个函数B。其返回值函数接受三个参数：
// context, request 以及 response interface.
// Token 值提取于consignment-ci中定义的上下文。我们将使用 user-service 认证这个 token.
// 如果认证通过，那么函数A会被执行，否则将返回错误。
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}
		// Note this is now uppercase (not entirely sure why this is...)
		token := meta["Token"]
		log.Println("Authenticating with token: ", token)
		// Auth here
		authClient := userService.NewUserServiceClient("go.micro.srv.user", client.DefaultClient)
		_, err := authClient.ValidateToken(context.Background(), &userService.Token{
			Token: token,
		})
		if err != nil {
			return err
		}
		err = fn(ctx, req, resp)
		return err
	}
}