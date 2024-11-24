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
	"path/filepath"
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
	router.Static("/uploads", "./uploads")

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
	auth.GET("/flights/:id/isBooked", h.IsFlightBookedByUser)
	auth.POST("/book", h.Book)
	auth.GET("/history", h.GetHistory)
	auth.POST("/cancel/:flightId", h.CancelBooking)

	admin := router.Group("/admin")
	admin.GET("/flights", h.GetFlights)
	admin.GET("/flights/:id", h.GetFlightById)
	admin.PUT("/flights/:id", h.UpdateFlight)
	admin.DELETE("/flights/:id", h.DeleteFlight)
	admin.POST("/flights", h.CreateFlight)

	api := router.Group("/api")
	api.GET("/airlinesaircrafts", h.GetAirlinesAircrafts)
	api.GET("/airports", h.GetAirports)

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

	user := types.UserLongData{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
		Phone:          req.Phone,
		DateOfBirth:    req.DateOfBirth,
		PassportSerie:  req.PassportSerie,
		PassportNumber: req.PassportNumber,
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
	ctx.Set("userImage", user.Image)
	ctx.Set("userIsAdmin", user.IsAdmin)
	ctx.JSON(http.StatusOK, gin.H{"token": token, "userImage": user.Image, "userIsAdmin": user.IsAdmin})
	return
}

func (h *Handler) Search(ctx *gin.Context) {
	dep := ctx.Query("dep")
	des := ctx.Query("des")
	depDate := ctx.Query("depDate")

	//log.Println(dep, des, depDate)
	if dep == "" || des == "" || depDate == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Can't get flight for this parameters"})
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
	//log.Println(ctx.Params)
	//log.Println(utils.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImV4YW1wbGVAZ21haWwuY29tIiwiZXhwIjoxNzI4Mzc3MTQ3fQ.xf9--F7bMBjOeSQZ4F7YTcv5vAoIR786BCA3O9tMHaU"))
	email, exists := ctx.Get("email")
	//log.Println(email)
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
		"image":          user.Image,
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

	isBooked, err := h.service.IsFlightBookedByUser(booking.FlightId, user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking booking status"})
		log.Println(err)
		return
	}

	if isBooked {
		ctx.JSON(http.StatusConflict, gin.H{"error": "You have already booked this flight"})
		return
	}

	log.Println(isBooked)

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

	var userData types.UserLongDataFromFront
	err := ctx.ShouldBind(&userData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("Bind err: ", err)
		return
	}

	if userData.Image != nil && userData.Image.Filename != "" {
		file := userData.Image
		filePath := filepath.Join("uploads", file.Filename)

		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
			log.Println("Failed to save image:", err)
			return
		}
	}

	userId, err := h.service.GetUserIdByEmail(fmt.Sprintf("%v", email))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}
	userData.Id = userId
	updateData, err := utils.ConvertToUserLongData(userData)
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

	status := ctx.Query("status")
	city := ctx.Query("city")

	var date *time.Time
	dateStr := ctx.Query("date")
	if dateStr != "" {
		parsedDate, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Expected YYYY-MM-DD."})
			return
		}
		date = &parsedDate
	}

	//log.Printf("email: %s, status: %s, city: %s, departureDate: %s", email, status, city, date)
	history, err := h.service.GetPassengerHistory(fmt.Sprintf("%v", email), status, city, date)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, history)
}

func (h *Handler) IsFlightBookedByUser(ctx *gin.Context) {
	flightId := ctx.Param("id")

	email, exists := ctx.Get("email")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, err := h.service.GetUserByEmail(email.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	flId, err := strconv.Atoi(flightId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Can't parse flight id"})
		return
	}
	isBooked, err := h.service.IsFlightBookedByUser(flId, user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not check booking status"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"isBooked": isBooked})
}

func (h *Handler) CancelBooking(ctx *gin.Context) {
	flightId := ctx.Param("flightId")

	flId, err := strconv.Atoi(flightId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Can't parse flight id"})
		return
	}

	email, exists := ctx.Get("email")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, err := h.service.GetUserByEmail(email.(string))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	err = h.service.CancelFlightByID(flId, user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Flight canceled successfully."})
	return
}

func (h *Handler) CreateFlight(ctx *gin.Context) {
	var flight types.FlightCreate
	if err := ctx.BindJSON(&flight); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateFlight(flight); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, flight)
}

func (h *Handler) GetFlights(ctx *gin.Context) {
	flights, err := h.service.GetAllFlights()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, flights)
}

func (h *Handler) GetFlightById(ctx *gin.Context) {
	flightId, _ := strconv.Atoi(ctx.Param("id"))

	flight, err := h.service.GetFlightById(flightId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, flight)
}

func (h *Handler) GetAirlinesAircrafts(ctx *gin.Context) {
	data, err := h.service.GetAirlinesAircrafts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (h *Handler) UpdateFlight(ctx *gin.Context) {
	var flight types.FlightControl
	if err := ctx.ShouldBindJSON(&flight); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(flight)
	if err := h.service.UpdateFlight(flight); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, flight)
}

func (h *Handler) DeleteFlight(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.DeleteFlight(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Flight deleted successfully"})
}

func (h *Handler) GetAirports(ctx *gin.Context) {
	airports, err := h.service.GetAirports()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, airports)
}
