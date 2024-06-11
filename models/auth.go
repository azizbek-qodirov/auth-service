package models

type RegisterReq struct {
	ID       string `json:"id"`       // User's unique identifier
	Username string `json:"username"` // User's username
	Email    string `json:"email"`    // User's email address
	Password string `json:"password"` // User's password
}

type RegisterResp struct {
	Message string `json:"message"` // Message indicating successful registration
}

type LoginReq struct {
	Username string `json:"username"` // User's username
	Password string `json:"password"` // User's password
}

type LoginResp struct {
	Token string `json:"token"` // JWT token provided after successful login
}

type GetProfileReq struct {
	Email string `json:"username"` // Username of the profile to retrieve
}

type GetProfileResp struct {
	ID       string `json:"id"`       // User's unique identifier
	Username string `json:"username"` // User's username
	Email    string `json:"email"`    // User's email address
	Password string `json:"password"` // User's password
}