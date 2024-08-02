package router

import (
	"payment_gateway/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/process_payment", controllers.ProcessPayment)
	r.GET("/retrieve_payment/:id", controllers.RetrievePayment)
	return r
}
