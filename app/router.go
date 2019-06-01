package app

import (
	"book-store/controller/Authentication"
	"book-store/controller/Book"
	"book-store/controller/Cart"
	"book-store/controller/Comment"
	"book-store/controller/Customer"
	"book-store/controller/Order"
	"book-store/controller/SetUp"

	"book-store/database"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

//SetupRouter ...
func SetupRouter() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	router.Use(cors.New(config))
	router.Use(static.Serve("/", static.LocalFile("./client", true)))

	api := router.Group("/api")

	api.POST("/sign_up", Authentication.SignUp)
	api.POST("/sign_in", Authentication.SignIn)

	auth := api.Group("/auth")
	{
		customer := auth.Group("/customer")
		{
			customer.GET("/:customer_id", Customer.GetCustomerByID)
			customer.GET("/:customer_id/book/:book_id/comments", Comment.GetCommentsByBookID)
			customer.GET("/:customer_id/books", Book.GetAllBooks)
			customer.GET("/:customer_id/book/:book_id", Book.GetBookInfoByID)
			customer.PUT("/:customer_id/customer/:customer_id", Customer.UpdateCustomerInfo)
			customer.POST("/:customer_id/cart", Cart.AddBookToCart)
			customer.PUT("/:customer_id/cart/:cart_id", Cart.UpdateCart)
			customer.GET("/:customer_id/carts", Cart.GetAllCartsInfo)
			customer.POST("/:customer_id/order", Order.CreateOrder)

			//auth.POST("/:customer_id/payment", Payment.AddPayment)
			//auth.DELETE("/:customer_id/payment", Payment.RemovePayment)

			//auth.POST("/:customer_id/shipping_info", ShippingInfo.CreateShippingInfo)
			//auth.PUT("/:customer_id/shipping_info", ShippingInfo.UpdateShippingInfo)
			//auth.DELETE("/:customer_id/shipping_info", ShippingInfo.DeleteShippingInfo)
		}
		staff := auth.Group("/staff")
		{
			staff.POST("/book", Book.CreateBook)
			staff.PUT("/book/:book_id", Book.UpdateBookInfo)
			staff.DELETE("/book/:book_id", Book.DeleteBook)
		}
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "404", "message": "Page not found"})
	})

	return router
}

//AppInit App Handler Start From Here
func AppInit() {
	SetUp.InitStuffs()
	database.InitMongo()
	router := SetupRouter()
	if err := router.Run(":2013"); err != nil {
		log.Fatal(err)
	}

}
