package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/proyuen/flashSale/Server/internal/db"
	"github.com/proyuen/flashSale/Server/internal/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	// 初始化 SQLC 的 Store (我们的数据库管家)
	store := db.New(conn)
	server := gin.Default()

	server.POST("/users", func(ctx *gin.Context) {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Email    string `json:"email"`
		}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		hashedPassword, err := util.HashPassword(req.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "password encryption failed"})
			return
		}
		arg := db.CreateUserParams{
			Username: req.Username,
			Password: hashedPassword,
			Email:    req.Email,
		}
		user, err := store.CreateUser(ctx, arg)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, user)
	})
	server.GET("/users/:username", func(ctx *gin.Context) {
		username := ctx.Param("username")
		user, err := store.GetUser(ctx, username)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, user)
	})
	log.Println("Server starting at", config.ServerAddress)
	err = server.Run(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
