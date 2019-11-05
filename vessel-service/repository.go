package main

import (
	pb "github.com/ChenHanZhang/microservices-in-golang/proto/vessel"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName = "shippy"
	vesselCollection = "vessels"
)

type Repository interface {
	FindAvailable(*pb.Specification)(*pb.Vessel, error)
	Create(*pb.Vessel) error
	Close()
}

type VesselRepository struct {
	session *mgo.Session
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	// 在之前我们使用了 vesselRepository 中一个 slice 来保存数据
	// 而这里我们改用了 db 来存储和访问数据
	var vessel *pb.Vessel

	err := repo.collection().Find(bson.M{
		"capacity": bson.M{ "$gte": spec.Capacity },
		"maxweight": bson.M{ "$gte": spec.MaxWeight },
	}).One(&vessel)

	if err != nil {
		return nil, err
	}

	return vessel, nil
}

func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}

func (repo *VesselRepository) Close() {
	repo.session.Close()
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(vesselCollection)
}