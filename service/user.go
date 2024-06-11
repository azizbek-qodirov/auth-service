package service

import (
	"auth-service/models"
	"auth-service/postgresql/managers"
	"database/sql"
)

type UserService struct {
	UM managers.UserManager
}

func NewUserService(conn *sql.DB) *UserService {
	return &UserService{UM: *managers.NewUserManager(conn)}
}

func (u *UserService) Register(req *models.RegisterReq) (*models.RegisterResp, error) {
	return nil, nil
}

func (u *UserService) Login(req *models.LoginReq) bool {
	return false

}

func (u *UserService) GetProfile(req *models.GetProfileReq) (*models.GetProfileResp, error) {
	return nil, nil
}
