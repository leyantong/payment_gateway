package router

import (
	"payment_gateway/controllers"
	"payment_gateway/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/process_payment", middleware.ValidatePaymentRequest(), controllers.ProcessPayment)
	r.GET("/retrieve_payment/:id", controllers.RetrievePayment)

	return r
}
