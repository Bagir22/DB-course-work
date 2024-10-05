package handler

import (
	"context"
	"courseWork/internal/service"
	"courseWork/internal/types"
	"courseWork/internal/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
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

	router.Use(cors.Default())

	router.POST("/signup", h.Signup)
	router.POST("/login", h.Login)
	router.GET("/search", h.Search)
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

	var req types.SignupRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	passportNumber, err := strconv.Atoi(req.PassportNumber)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid passport number"})
		return
	}

	user := types.UserLongData{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		Phone:          req.Phone,
		DateOfBirth:    req.DateOfBirth,
		PassportSerie:  req.PassportSerie,
		PassportNumber: passportNumber,
		Password:       req.Password,
	}

	userResponse, err := h.service.AddUser(context.TODO(), user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, types.Response{"Can't save user", err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, userResponse)
	return
}

func (h *Handler) Login(ctx *gin.Context) {
	/*
		{
			"email": "example@gmail.com",
			"password": "SomePassword"
		}
	*/

	var user types.UserShortData
	if err := ctx.BindJSON(&user); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user, err := h.service.CheckUserExist(user.Email, user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	token, err := utils.GenerateToken(user.Email)
	if err != nil {
		log.Println("Error generating token:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) Search(ctx *gin.Context) {
	dep := ctx.Query("dep")
	des := ctx.Query("des")
	depDate := ctx.Query("depDate")

	log.Println(dep, des, depDate)
	if dep == "" || des == "" || depDate == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "1 Can't get flight for this parameters"})
		return
	}

	departureDate, err := time.Parse("2006-01-02", depDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format; must be YYYY-MM-DD"})
		return
	}

	flights, err := h.service.GetFlights(dep, des, departureDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "2 Can't get flight for this parameters"})
		return
	}

	ctx.JSON(http.StatusOK, flights)
}
