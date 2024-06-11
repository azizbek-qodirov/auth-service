package managers

import (
	"auth-service/models"
	"database/sql"
)

type UserManager struct {
	Conn *sql.DB
}

func NewUserManager(db *sql.DB) *UserManager {
	return &UserManager{Conn: db}
}

func (m *UserManager) Register(req models.RegisterReq) (*models.RegisterResp, error) {
	return nil, nil
}

func (m *UserManager) Login(req models.LoginReq) (*models.LoginResp, error) {
	return nil, nil
}

func (m *UserManager) Profile(req models.GetProfileReq) (*models.GetProfileResp, error) {
	return nil, nil
}
