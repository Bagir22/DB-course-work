package main

import (
	"courseWork/internal/config"
	"courseWork/internal/handler"
	"courseWork/internal/server"
	"courseWork/internal/service"
	"courseWork/internal/storage/postgres"
	"courseWork/internal/utils"
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

	err = utils.MarkFlightsAsDone(conn.DB)
	if err != nil {
		log.Println("Error marking flights as done:", err)
	}

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

/*
require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/cors v1.7.2
	github.com/gin-gonic/gin v1.10.0
	github.com/jmoiron/sqlx v1.4.0
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.9
	golang.org/x/crypto v0.28.0
)

*/
