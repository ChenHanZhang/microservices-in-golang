package main

import (
	"context"
	pb "github.com/ChenHanZhang/microservices-in-golang-proto/user"
	micro "github.com/micro/go-micro"
	"log"
)

const topic = "user.created"

type Subscriber struct {
}

func (sub *Subscriber) Process(ctx context.Context, user *pb.User) error {
	log.Println("Picked up a new message")
	log.Println("Sending email to:", user.Name)
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.email"),
		micro.Version("latest"))
	srv.Init()

	pubSub := srv.Server().Options().Broker

	if err := pubSub.Connect(); err != nil {
		log.Fatalf("broker connect error: %v\n", err)
	}

	// 订阅消息
	err := micro.RegisterSubscriber(topic, srv.Server(), new(Subscriber))

	if err = srv.Run(); err != nil {
		log.Fatal(err)
	}
}
