package handlers

import (
	t "api-gateway/api/token"
	"api-gateway/service"
	s "api-gateway/storage"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)



type UserReq struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

// @Router 				/register/user [post]
// @Summary 			REGISTER USER
// @Description		 	This api registers user
// @Tags 				USER
// @Accept 				json
// @Produce 			json
// @Param data 			body UserReq true "UserReq"
// @Success 201 		{object} models.UserReq
// @Failure 400 		string Error
func (h *Handler) Register(c *gin.Context) {
	var (
		req UserReq
	)

	err := c.ShouldBindJSON(&req)
	log.Println()
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": "invalid data",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, req)

}

// @Router 				/login [post]
// @Summary 			Login USER
// @Description		 	This api registers user
// @Tags 				USER
// @Accept 				json
// @Produce 			json
// @Param data 			body models.LoginReq true "LoginReq"
// @Success 201 		{object} token.Tokens
// @Failure 400 		string Error
func (h *Handler) Login(c *gin.Context) {
	var req models.LoginReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.IndentedJSON(400, gin.H{
			"error": err.Error(),
		})
	}

	user, err := s.GetUserByUsername(req.Username)
	if err != nil {
		c.IndentedJSON(http.StatusForbidden, gin.H{
			"error": "user not found",
		})
		return
	}

	if req.Password == user.Passsword && req.Username == user.Username {
		tokens := t.GenereteJWTToken(user.Id, user.UserRole, user.Username)
		c.IndentedJSON(200, tokens)
		return
	}

	c.IndentedJSON(http.StatusForbidden, gin.H{
		"error": "incorrect username or password",
	})

}

// @Router 				/user [get]
// @Summary 			GET USER
// @Description		 	This api GETS user by id
// @Tags 				USER
// @Accept 				json
// @Produce 			json
// @Security    		BearerAuth
// @Param 				id query string true "ID"
// @Success 200			{object} models.UserReq
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *Handler) GetUser(c *gin.Context) {

	id := c.Query("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "invalid ID",
		})
		return
	}

	if intId != 7 {
		c.IndentedJSON(404, gin.H{
			"error": "user not found",
		})
		return
	}

	res := UserReq{
		Name:     "Azizbek",
		Username: "Sobirov",
		Password: "12345677",
		Age:      17,
	}

	c.IndentedJSON(http.StatusOK, res)
}

// @Router 				/users [get]
// @Summary 			GET USERS
// @Description		 	This api GETS allusers
// @Tags 				USER
// @Accept 				json
// @Produce 			json
// @Success 200			{object} models.UserReq
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *Handler) GetUsers(c *gin.Context) {

	res := []UserReq{
		{
			Name:     "Azizbek",
			Username: "Sobirov",
			Password: "12345677",
			Age:      17,
		},
		{
			Name:     "Dilshodbek",
			Username: "umarov",
			Password: "0000",
			Age:      19,
		},
	}

	c.IndentedJSON(http.StatusOK, res)
}
