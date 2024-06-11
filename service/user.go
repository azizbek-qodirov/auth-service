package service

import (
	pb "auth-service/RestaurantRservationSubmodule/genprotos"
	"auth-service/postgresql/managers"
	"database/sql"
)

type UserService struct {
	UM managers.UserManager
}

func NewUserService(conn *sql.DB) *UserService {
	return &UserService{UM: *managers.NewUserManager(conn)}
}

func (u *UserService) Register(req *pb.RegisterReq) (*pb.RegisterResp, error) {
	return nil, nil
}

func (u *UserService) GetUser(id int) (*m.UserRes, error) {
	return nil, nil
}

func (u *UserService) Login(req *m.LoginReq) bool {
	return false

}
