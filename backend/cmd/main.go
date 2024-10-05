package main

import (
	"courseWork/internal/config"
	"courseWork/internal/handler"
	"courseWork/internal/server"
	"courseWork/internal/service"
	"courseWork/internal/storage/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := godotenv.Load()

	cfg := config.InitConfig()

	conn, err := postgres.InitConn(cfg)
	if err != nil {
		log.Println(err)
		log.Fatal("Can't init connection to database")
	}
	defer conn.Close()

	db := postgres.InitDb(conn)

	service := service.InitService(db)

	handler := handler.InitHandler(service)

	server.Run(handler)

	c := make(chan os.Signal, 1)
	signal.Notify(c,
		syscall.SIGTERM,
		syscall.SIGINT,
	)

	<-c
}
