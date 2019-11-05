package main

import (
	pb "github.com/ChenHanZhang/microservices-in-golang/proto/consignment"

	"context"
	"encoding/json"
	"errors"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config/cmd"
	"io/ioutil"
	"log"
	"os"
)

const (
	ADDRESS = "localhost:50051"
	DEFAULT_INFO_FILE = "consignment.json"
)

// 读取 consignment.json 中的货物信息
func parseFile(fileName string) (*pb.Consignment, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var consignment *pb.Consignment
	err = json.Unmarshal(data, &consignment)
	if err != nil {
		return nil, errors.New("consignment.json file content error")
	}
	return consignment, nil
}

func main() {

	err := cmd.Init()

	// 服务发现 => pb.NewShippingService TODO: 这个方法封装了什么？
	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)

	// 可以在命令行中指定 json 文件
	infoFile := DEFAULT_INFO_FILE
	if len(os.Args) > 1 {
		infoFile = os.Args[1]
	}

	// 解析货物
	consignment, err := parseFile(infoFile)
	if err != nil {
		log.Fatalf("parse info file error: %v", err)
	}

	// 调用 RPC
	// 将货物存储到自己的仓库里
	// TODO: 这样就实现了如同本地调用函数一样的方式来调用远程方法
	//
	resp, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("create consignment error: %v", err)
	}

	log.Printf("created: %t", resp.Created)

	// 列出当前所有托运的货物
	resp, err = client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("failed to list consignments: %v", err)
	}
	for _, c := range resp.Consignments{
		log.Printf("%+v", c)
	}
}