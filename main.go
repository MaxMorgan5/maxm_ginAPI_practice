package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

var products = []product{
	{ID: "1", Name: "Oreos", Price: 475, Quantity: 400},
	{ID: "2", Name: "Milk", Price: 350, Quantity: 400},
	{ID: "3", Name: "Orange", Price: 119, Quantity: 390},
	{ID: "4", Name: "Bread", Price: 275, Quantity: 211},
}

func getProducts(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, products)
}

func main() {
	router := gin.Default()
	router.GET("/products", getProducts)

	router.Run("localhost:9090")

}
