package api

import "github.com/gin-gonic/gin"

// Server serves HTTP requests for the api service
type Server struct {
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer() *Server {

	server := &Server{}

	router := gin.Default()

	router.POST("/users", server.createUser)

	server.router = router
	return server

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
