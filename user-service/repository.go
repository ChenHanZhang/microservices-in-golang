package main

import (
	pb "github.com/ChenHanZhang/microservices-in-golang-proto/user"

	"github.com/jinzhu/gorm"
)

type Repostory interface {
	GetAll() ([]*pb.User, error)
	Get(id string) (*pb.User, error)
	Create(user *pb.User) error
	GetByEmailAndPassword(user *pb.User) (*pb.User, error)
}

// 要习惯这种写法：
// 先根据 proto 规定好接口
// 然后声明一个 struct
// 该 struct 用以实现上述接口
// 这样有个好处，就是数据与实现解耦
type UserRepostory struct {
	db *gorm.DB
}

func (repo *UserRepostory) Get(id string) (*pb.User, error) {
	var user *pb.User
	user.Id = id
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepostory) Create(user *pb.User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepostory) GetByEmailAndPassword(user *pb.User) (*pb.User, error) {
	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepostory) GetAll() ([]*pb.User, error)  {
	var users []*pb.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

