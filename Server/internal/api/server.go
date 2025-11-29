package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/proyuen/flashSale/Server/internal/db"
	"github.com/proyuen/flashSale/Server/internal/token"
	"github.com/proyuen/flashSale/Server/internal/util"
)

type Server struct {
	store      *db.Queries
	router     *gin.Engine
	config     util.Config
	tokenMaker token.Maker
}

func NewServer(config util.Config, store *db.Queries) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		store:      store,
		config:     config,
		tokenMaker: tokenMaker,
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
