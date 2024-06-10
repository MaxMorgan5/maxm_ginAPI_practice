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

// indented json serializes product struct objects into json
func getProducts(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, products)
}

// BindJson uses the json tags in our struct to map the json we're sending in body of our request, to the fields in our struct.
// func makes new product var, we get the json from the body of the request, pass a reference to the new var and bind json to struct format
// append to in mem products slice, and return http response code and json of newly created product
func addProduct(context *gin.Context) {
	var newProduct product

	if err := context.BindJSON(&newProduct); err != nil {
		return
	}

	products = append(products, newProduct)

	context.IndentedJSON(http.StatusCreated, newProduct)
}

func main() {
	router := gin.Default()
	router.GET("/products", getProducts)

	router.Run("localhost:9090")

}
