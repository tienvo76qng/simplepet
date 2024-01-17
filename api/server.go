package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/tienvo76qng/simplepet/db/sqlc"
)

// Server server HTTP requests for our banking service
type Server struct {
	store  *db.Store
	router *gin.Engine // help us send each API request to the correct handler for processing
}

// NewServer creates a new Http server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// add routes to router
	router.POST("/accounts", server.CreateAccount)
	router.GET("/accounts/:id", server.GetAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
