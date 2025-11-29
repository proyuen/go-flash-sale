package api

import (
	"github.com/gin-gonic/gin"
	"github.com/proyuen/flashSale/Server/internal/db"
)

type Server struct {
	store  *db.Queries
	router *gin.Engine
}

func NewServer(store *db.Queries) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/users", server.createUesr)
	router.GET("/users/:username", server.getUser)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
