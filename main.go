package main

import (
	nykeProducts "proboard/api/src/nyke/products"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/nyke/dashboard_product", nykeProducts.GetDashboardProduct)

	router.Run("localhost:8080")
}
