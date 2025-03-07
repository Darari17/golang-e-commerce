package router

import (
	"log"

	"github.com/Darari17/golang-e-commerce/config"
	"github.com/Darari17/golang-e-commerce/handler"
	"github.com/Darari17/golang-e-commerce/middleware"
	"github.com/Darari17/golang-e-commerce/migrate"
	"github.com/Darari17/golang-e-commerce/repository"
	"github.com/Darari17/golang-e-commerce/security"
	"github.com/Darari17/golang-e-commerce/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	authHandler        handler.IAuthHandler
	userHandler        handler.IUserHandler
	productHandler     handler.IProductHandler
	transactionHandler handler.ITransactionHandler

	authMiddleware middleware.IAuthMiddleware
	jwtHandler     security.IJWTHandler

	app *gin.Engine
	db  *gorm.DB
}

func (s *Server) setupRoutes() {
	api := s.app.Group("/api/v1")

	api.POST("/auth/login", s.authHandler.Login)
	api.POST("/register", s.userHandler.Register)
	api.GET("/profile", s.authMiddleware.RequiredToken(), s.userHandler.Profile)

	product := api.Group("/products")
	{
		product.POST("/", s.authMiddleware.RequiredToken("admin"), s.productHandler.CreateProduct)
		product.PUT("/:id", s.authMiddleware.RequiredToken("admin"), s.productHandler.UpdateProduct)
		product.DELETE("/:id", s.authMiddleware.RequiredToken("admin"), s.productHandler.DeleteProduct)
		product.GET("/:id", s.productHandler.FindProductById)
		product.GET("/", s.productHandler.FindAllProducts)
	}

	tx := api.Group("/transactions")
	{
		tx.POST("/", s.authMiddleware.RequiredToken("user"), s.transactionHandler.CreateTransaction)
		tx.GET("/", s.authMiddleware.RequiredToken("user"), s.transactionHandler.FindAllTransactionsByUserId)
		tx.GET("/:id", s.authMiddleware.RequiredToken("admin", "user"), s.transactionHandler.FindTransactionById)
		tx.PATCH("/:id/cancel", s.authMiddleware.RequiredToken("user"), s.transactionHandler.CancelTransaction)
		tx.PATCH("/:id/status", s.authMiddleware.RequiredToken("admin"), s.transactionHandler.UpdateTransactionStatus)
		tx.DELETE("/:id", s.authMiddleware.RequiredToken("admin"), s.transactionHandler.DeleteTransaction)
	}
}

func (s *Server) Run(port string) {
	defer func() {
		sqlDB, err := s.db.DB()
		if err != nil {
			log.Println("Failed to get database instance:", err)
			return
		}

		err = sqlDB.Close()
		if err != nil {
			log.Println("Failed to close database connection:", err)
			return
		}

		log.Println("Database connection closed successfully")
	}()

	s.setupRoutes()
	log.Println("Server is running on port", port)

	err := s.app.Run(port)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func NewServer() *Server {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	err = migrate.AutoMigrate(db)
	if err != nil {
		log.Fatal("Database migration failed:", err)
	}

	app := gin.Default()
	jwtHandler := security.NewJwtHandler("kunciRahasia") //seharusnya tidak hardcode
	authMiddleware := middleware.NewAuthMiddleware(jwtHandler)

	userRepository := repository.NewUserRepository(db)
	productRepository := repository.NewProductRepository(db)
	orderRepository := repository.NewOrderRepository()
	orderItemRepository := repository.NewOrderItemRepository()

	authService := service.NewAuthService(userRepository, jwtHandler)
	userService := service.NewUserService(userRepository)
	productService := service.NewProductRepository(productRepository)
	transactionService := service.NewTransactionService(orderRepository, orderItemRepository, productRepository, db)

	authHandler := handler.NewAuthService(authService)
	userHandler := handler.NewUserHandler(userService)
	productHandler := handler.NewProducyHandler(productService)
	transactionHandler := handler.NewTxHandler(transactionService)

	return &Server{
		authHandler:        authHandler,
		userHandler:        userHandler,
		productHandler:     productHandler,
		transactionHandler: transactionHandler,
		authMiddleware:     authMiddleware,
		jwtHandler:         jwtHandler,
		app:                app,
		db:                 db,
	}
}
