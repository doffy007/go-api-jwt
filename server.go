package main

import (
	"github.com/doffy007/go-api-jwt/config"
	"github.com/doffy007/go-api-jwt/controller"
	"github.com/doffy007/go-api-jwt/middleware"
	"github.com/doffy007/go-api-jwt/repository"
	"github.com/doffy007/go-api-jwt/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                *gorm.DB                     = config.SetupDatabaseConnection()
	userRepository    repository.UserRepository    = repository.NewUserRepository(db)
	productRepository repository.ProductRepository = repository.NewProductRepository(db)
	jwtService        service.JWTService           = service.NewJWTService()
	userService       service.UserService          = service.NewUserservice(userRepository)
	productService    service.ProductService       = service.NewProductService(productRepository)
	authService       service.AuthService          = service.NewAuthService(userRepository)
	authController    controller.AuthController    = controller.NewAuthController(authService, jwtService)
	userController    controller.UserController    = controller.NewUserController(userService, jwtService)
	productController controller.ProductController = controller.NewProductController(productService, jwtService)
)

//Get main page
func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}


	productRoutes := r.Group("api/products", middleware.AuthorizeJWT(jwtService))
	{
		productRoutes.GET("/", productController.All)
		productRoutes.POST("/", productController.Insert)
		productRoutes.GET("/:id", productController.FindByID)
		productRoutes.PUT("/:id", productController.Update)
		productRoutes.DELETE("/:id", productController.Delete)
	}
	
	r.Run(":3306")
}
