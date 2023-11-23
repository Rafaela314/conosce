package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Name string `json:"name" binding:"required"`
}

var users []string

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// TODO: this simulates DB tx behavior. Add db driver tx latter on
	users = append(users, req.Name)

	ctx.JSON(http.StatusOK, req.Name)

}
