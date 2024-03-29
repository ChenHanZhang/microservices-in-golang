package main

import (
	pb "github.com/ChenHanZhang/microservices-in-golang-proto/user"
	"github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	"golang.org/x/net/context"
	"log"
	"os"
)


func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.user-cli"),
		micro.Version("latest"),
	)
	// Init will parse the command line flags.
	srv.Init()
	client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)
	name := "tfhnia"
	email := "tfhnia@163.com"
	password := "123456"
	company := "BUPT"
	r, err := client.Create(context.TODO(), &pb.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %s", r.User.Id)
	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}
	for _, v := range getAll.Users {
		log.Println(v)
	}
	authResponse, err := client.Auth(context.TODO(), &pb.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Fatalf("Could not authenticate user: %s error: %v\n", email, err)
	}
	log.Printf("Your access token is: %s \n", authResponse.Token)
	// let's just exit because
	os.Exit(0)
}