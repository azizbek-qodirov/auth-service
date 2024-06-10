package service

import (
	m "api-gateway/models"
	"fmt"
	"log"
	"net/rpc"
)

type UserService struct {
	Client *rpc.Client
}

func (u *UserService) Register(req *m.UserReq) (*m.UserRes, error) {

	res := m.UserRes{}

	err := u.Client.Call("User.Register", req, &res)

	if err != nil {
		log.Fatal("Client invocation error: ", err)
		return nil, err
	}

	fmt.Println("User Registerd : ", res)

	return &res, nil

}

func (u *UserService) GetUser(id int) (*m.UserRes, error) {

	res := m.UserRes{}
	err := u.Client.Call("User.GetUserData", id, &res)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Printf("User with %d id : %v\n", res.Id, res)

	return &res, nil

}

func (u *UserService) Login(req *m.LoginReq) bool {

	res := new(m.LoginRes)

	fmt.Println(req)

	err := u.Client.Call("User.Login", req, &res)
	if err != nil {
		log.Println(err)
		return false
	}

	if res.Correct {
		return true
	}

	return false

}
