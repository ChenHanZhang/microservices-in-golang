package main

import (
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
	"log"

	pb "github.com/ChenHanZhang/microservices-in-golang-proto/consignment"
	vesselProto "github.com/ChenHanZhang/microservices-in-golang-proto/vessel"
)

type service struct {
	session *mgo.Session
	vesselClient vesselProto.VesselServiceClient
}

// 请注意session.Clone()，为什么要Clone?
func (s *service) GetRepo() Repository {
	return &ConsignmentRepository{s.session.Clone()}
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	log.Println("creating consignment...")
	log.Println("Looking for a available vessel..")

	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity: int32(len(req.Containers)),
	})
	if err != nil {
		// TODO: 这里有必要将错误给 Wrap 上吗
		log.Printf("vesselClient cause an error: %v\n", err)
		return err
	}
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	// We set the VesselId as the vessel we got back from our
	// vessel service
	req.VesselId = vesselResponse.Vessel.Id
	// Save our consignment
	err = repo.Create(req)
	if err != nil {
		return err
	}
	// Return matching the `Response` message we created in our
	// protobuf definition.
	res.Created = true
	res.Consignment = req
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()
	consignments, err := repo.GetAll()
	if err != nil {
		return err
	}
	res.Consignments = consignments
	return nil
}