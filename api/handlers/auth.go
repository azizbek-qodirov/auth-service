package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h *HTTPHandler) Register(c *gin.Context) {

}

func (h *HTTPHandler) Login(c *gin.Context) {

}

// GetProfile godoc
// @ID getprofile
// @Router /profile [GET]
// @Summary Get Profile
// @Description Get Profile
// @Tags party
// @Accept json
// @Produce json
// @Param id path string true "party ID"
// @Success 200 {object} pb.PartyGetById "party data"
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "party not found"
// @Failure 500 {object} string "Server error"
func (h *HTTPHandler) Profile(c *gin.Context) {

}
