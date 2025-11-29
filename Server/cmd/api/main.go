package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/proyuen/flashSale/Server/internal/api"
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
	server := api.NewServer(store)

	log.Println("Server starting at", config.ServerAddress)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
