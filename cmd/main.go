package main

import (
	r "api-gateway/api"
	h "api-gateway/api/handler"

	"api-gateway/service"
	"log"
)

func main() {

	userService := service.UserService{}

	userHandler := h.NewHandler(&userService)

	router := r.Router(userHandler)

	log.Println("api-gateway run in :7070 ")
	if err := router.Run(":7071"); err != nil {
		log.Fatal(err)
	}

}
