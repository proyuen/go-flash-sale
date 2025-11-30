package api

import (
	"github.com/gin-gonic/gin"
	"github.com/proyuen/flashSale/Server/internal/service"
	"github.com/proyuen/flashSale/Server/internal/util"
)

type Server struct {
	service service.Service
	router  *gin.Engine
	config  util.Config
}

func NewServer(config util.Config, service service.Service) (*Server, error) {
	server := &Server{
		service: service,
		config:  config,
	}
	router := gin.Default()

	router.POST("/users", server.createUesr)
	router.POST("/users/login", server.loginUser)
	router.GET("/users/:username", server.getUser)

	server.router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
