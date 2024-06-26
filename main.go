package main

import (
	"errors"
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

// helper func for router func. Loops over products slice to search for ID that matches id in arg
// returns a pointer to the correct product or an err if not found
func productIDMatch(id string) (*product, error) {
	for index, product := range products {
		if product.ID == id {
			return &products[index], nil
		}
	}

	return nil, errors.New("product not found")
}

// retrieve ID from http request.
// call helper func and store result, check for err.
// if no err, return the matching product as a json response and a 200 status code
func getProductByID(context *gin.Context) {
	id := context.Param("id")

	matchingProduct, err := productIDMatch(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, matchingProduct)
}

func updateQuantity(context *gin.Context) {
	id := context.Param("id")

	matchingProduct, err := productIDMatch(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}

	if err := context.BindJSON(&matchingProduct); err != nil {
		return
	}

	context.IndentedJSON(http.StatusOK, matchingProduct)

}

func main() {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProductByID)
	router.PATCH("/products/:id", updateQuantity)
	router.POST("/products", addProduct)
	router.Run("localhost:9090")

}
