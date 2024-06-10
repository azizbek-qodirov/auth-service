package handlers

import "auth-service/service"

type HTTPHandler struct {
}

func NewHandler(userService *service.UserService) *HTTPHandler {
	return &HTTPHandler{}
}
