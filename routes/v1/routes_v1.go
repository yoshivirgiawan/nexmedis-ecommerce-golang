package v1

import (
	v1Controller "ecommerce/app/http/controllers/v1"
	"ecommerce/app/http/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		authController := v1Controller.NewAuthController()
		userController := v1Controller.NewUserController()
		productController := v1Controller.NewProductController()
		cartController := v1Controller.NewCartController()
		orderController := v1Controller.NewOrderController()

		authRoute := v1.Group("/auth")
		{
			authRoute.POST("/register", authController.Register)
			authRoute.POST("/login", authController.Login)
			authRoute.POST("/refresh_token", authController.RefreshToken)
			authRoute.GET("/fetch", middlewares.AuthMiddleware(), authController.FetchUser)
		}

		usersRoute := v1.Group("/users")
		{
			usersRoute.POST("/email_checkers", userController.CheckEmailAvailability)
			usersRoute.GET("/", middlewares.AuthMiddleware("admin"), userController.GetUsers)
			usersRoute.POST("/", middlewares.AuthMiddleware("admin"), userController.Create)
			usersRoute.GET("/:id", middlewares.AuthMiddleware("admin"), userController.GetUser)
			usersRoute.PUT("/:id", middlewares.AuthMiddleware("admin"), userController.Update)
			usersRoute.DELETE("/:id", middlewares.AuthMiddleware("admin"), userController.Delete)
		}

		productsRoute := v1.Group("/products")
		{
			productsRoute.GET("/", productController.GetProducts)
			productsRoute.POST("/", middlewares.AuthMiddleware("admin"), productController.Create)
			productsRoute.GET("/:id", productController.GetProduct)
			productsRoute.PUT("/:id", middlewares.AuthMiddleware("admin"), productController.Update)
			productsRoute.DELETE("/:id", middlewares.AuthMiddleware("admin"), productController.Delete)
		}

		cartsRoute := v1.Group("/carts")
		{
			cartsRoute.GET("/", middlewares.AuthMiddleware("user"), cartController.GetCarts)
			cartsRoute.POST("/", middlewares.AuthMiddleware("user"), cartController.AddToCart)
			cartsRoute.GET("/:id", middlewares.AuthMiddleware("user"), cartController.GetCart)
			cartsRoute.PUT("/:id", middlewares.AuthMiddleware("user"), cartController.Update)
			cartsRoute.DELETE("/:id", middlewares.AuthMiddleware("user"), cartController.Delete)
		}

		ordersRoute := v1.Group("/orders")
		{
			ordersRoute.GET("/", middlewares.AuthMiddleware("user"), orderController.GetOrders)
			ordersRoute.POST("/", middlewares.AuthMiddleware("user"), orderController.Create)
			ordersRoute.GET("/:id", middlewares.AuthMiddleware("user"), orderController.GetOrder)
			ordersRoute.GET("/:id/pay", middlewares.AuthMiddleware("user"), orderController.PayOrder)
		}
	}
}
