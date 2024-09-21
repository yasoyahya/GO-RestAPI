package main

import (
	"GO-RESTAPI-GIN/controllers/categorycontroller"
	"GO-RESTAPI-GIN/controllers/productcontroller"
	"GO-RESTAPI-GIN/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	// Product
	r.GET("/api/product", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Delete)

	// Category
	r.GET("/api/category", categorycontroller.Index)
	r.GET("/api/category/:id", categorycontroller.Show)
	r.POST("/api/category", categorycontroller.Create)
	r.PUT("/api/category/:id", categorycontroller.Update)
	r.DELETE("/api/category", categorycontroller.Delete)

	r.Run()
}
