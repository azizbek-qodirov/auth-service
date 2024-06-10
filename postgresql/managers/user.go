package managers

import (
	"auth-service/RestaurantRservationSubmodule/genprotos"
	"database/sql"
)

type UserManager struct {
	Conn *sql.DB
}

func NewUserManager(db *sql.DB) *UserManager {
	return &UserManager{Conn: db}
}

func (m *UserManager) Register(req genprotos.RegisterReq) (*genprotos.RegisterResp, error) {
	
}

func (m *UserManager) Login(req genprotos.LoginReq) (*genprotos.LoginResp, error) {

}

func (m *UserManager) Profile(req genprotos.GetProfileReq) (*genprotos.GetProfileResp, error) {

}
