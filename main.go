package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Products []ProductItem
}

type ProductItem struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    Category_Type
}

type Category_Type struct {
	ID           string `json:"id"`
	CategoryName string `json:"categoryName"`
}

func main() {
	router := gin.Default()
	router.GET("/dashboard_product", getDashboardProduct)

	router.Run("localhost:8080")
}

func getDashboardProduct(c *gin.Context) {

	content, err := ioutil.ReadFile("./mock_data/product_dashboard.json")

	// check if file has an error
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error on json file"})
	}

	fmt.Println(string(content))
	var payload Product
	err = json.Unmarshal(content, &payload)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
	}

	c.IndentedJSON(http.StatusOK, payload.Products)
}
