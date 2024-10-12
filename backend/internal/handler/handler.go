package handler

import (
	"context"
	"courseWork/internal/middleware"
	"courseWork/internal/service"
	"courseWork/internal/types"
	"courseWork/internal/utils"
	"fmt"
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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.POST("/signup", h.Signup)
	router.POST("/login", h.Login)
	router.GET("/search", h.Search)

	auth := router.Group("/api")
	auth.Use(middleware.AuthMiddleware())

	auth.GET("/user", h.GetUser)
	auth.PUT("/user", h.UpdateUser)
	auth.GET("/flights/:id/seats", h.GetSeatsForFlight)
	auth.POST("/book", h.Book)
	auth.GET("/history", h.GetHistory)

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

	log.Println(token)
	ctx.JSON(http.StatusOK, gin.H{"token": token})
	return
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
		log.Println(flights, err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, flights)
	return
}

func (h *Handler) GetUser(ctx *gin.Context) {
	log.Println(ctx.Params)
	//log.Println(utils.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImV4YW1wbGVAZ21haWwuY29tIiwiZXhwIjoxNzI4Mzc3MTQ3fQ.xf9--F7bMBjOeSQZ4F7YTcv5vAoIR786BCA3O9tMHaU"))
	email, exists := ctx.Get("email")
	log.Println(email)
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, err := h.service.GetUserByEmail(email.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"firstName":      user.FirstName,
		"lastName":       user.LastName,
		"email":          user.Email,
		"phone":          user.Phone,
		"dateOfBirth":    user.DateOfBirth,
		"passportSerie":  user.PassportSerie,
		"passportNumber": user.PassportNumber,
	})
	return
}

func (h *Handler) GetSeatsForFlight(ctx *gin.Context) {
	flightId := ctx.Param("id")

	id, err := strconv.Atoi(flightId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	seats, err := h.service.GetSeatsForFlight(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, seats)
	return
}

func (h *Handler) Book(ctx *gin.Context) {
	log.Println(ctx.Params)
	var booking types.BookFlight
	err := ctx.BindJSON(&booking)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	email, exists := ctx.Get("email")
	log.Println(email)
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, err := h.service.GetUserByEmail(email.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	booking.PassengerId = user.Id

	log.Println(booking)
	err = h.service.AddFlightBooking(booking)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, booking)
	return
}

func (h *Handler) UpdateUser(ctx *gin.Context) {
	email, exists := ctx.Get("email")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var updateData types.UserLongData
	err := ctx.BindJSON(&updateData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	userId, err := h.service.GetUserIdByEmail(fmt.Sprintf("%v", email))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println(err)
	}

	updateData.Id = userId

	err = h.service.UpdateUser(updateData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, updateData)
	return
}

func (h *Handler) GetHistory(ctx *gin.Context) {
	email, exists := ctx.Get("email")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	history, err := h.service.GetPassengerHistory(fmt.Sprintf("%v", email))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, history)
	return
}
