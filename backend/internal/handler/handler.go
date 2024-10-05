package handler

import (
	"context"
	"courseWork/internal/service"
	"courseWork/internal/types"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Handler struct {
	service service.Repository
}

func InitHandler(service service.Repository) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()
	router.POST("/signup", h.Signup)
	return router
}

func (h *Handler) Signup(ctx *gin.Context) {
	/*
		{
		    "firstName": "SomeName",
		    "lastName": "SomeSurname",
		    "email": "example@gmail.com",
		    "phone": 89021009988,
		    "dateOfBirth": "2019-05-17",
		    "passportSerie": "ABC 1234",
		    "passportNumber": 123456,
		    "password": "SomePassword"
		}
	*/

	var user types.UserLongData
	err := ctx.BindJSON(&user)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, types.Response{"Can't parse user", err.Error()})
		return
	}

	userResponse, err := h.service.AddUser(context.TODO(), user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.Response{"Can't save user", err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, userResponse)
	return
}
