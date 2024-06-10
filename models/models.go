package models

type UserReq struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Users struct {
	Allusers []Userr
}

type Userr struct {
}

type LoginRes struct {
	Correct bool
}

type UserRes struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	UserRole  string `json:"user_role"`
	Age       int    `json:"age"`
	Passsword string `json:"password"`
}
