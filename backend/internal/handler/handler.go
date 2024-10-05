package handler

import (
	"courseWork/internal/service"
	"github.com/gin-gonic/gin"
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

	return router
}
