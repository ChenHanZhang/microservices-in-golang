package main

import (
	"fmt"
	pb "github.com/ChenHanZhang/microservices-in-golang-proto/user"
	"github.com/micro/go-micro"
	"log"
)

func main() {

	db, err := CreateConnection()
	if err != nil{
		log.Fatalf("Could not coonect to DB: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&pb.User{})

	repo := &UserRepostory{db}

	tokenService := &TokenService{repo}

	srv := micro.NewService(
			micro.Name("go.micro.srv.user"),
			micro.Version("latest"),
		)

	// 会自动读取命令行参数
	srv.Init()

	publisher := micro.NewPublisher(topic, srv.Client())

	pb.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService, publisher})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}