package api

import (
	"github.com/gin-gonic/gin"
	database "github.com/oluwaferanmiadetunji/simple_bank/internal/database"
)

type Server struct {
	store  *database.Store
	router *gin.Engine
}

func NewServer(store *database.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	server.router = router

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
