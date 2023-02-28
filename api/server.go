package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/igua95/simplebank/db/sqlc"
)

type Server struct {
	store  db.Store // of type db.Store
	router *gin.Engine
}

func NewService(store db.Store) *Server {
	server := &Server{store: store}

	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	router.DELETE("/accounts/:id", server.deleteAccount)

	server.router = router

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// gin.H store key value pairs
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
