package api

import (
	"fmt"

	"github.com/caard0s0/united-atomic-bank-server/configs"
	db "github.com/caard0s0/united-atomic-bank-server/internal/database/sqlc"
	"github.com/caard0s0/united-atomic-bank-server/pkg/token"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Server serves HTTP requests for our banking services
type Server struct {
	config     configs.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(config configs.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.Use(CORSMiddleware(server.config.HttpClientAddress))

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/transfers", server.createTransfer)
	authRoutes.GET("/transfers", server.listTransfers)
	authRoutes.POST("/loans", server.createLoanTransfer)

	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccounts)

	server.router = router
}

func CORSMiddleware(clientAddress string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", clientAddress)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Set-Cookie")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(serverAddress string) error {
	return server.router.Run(serverAddress)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
