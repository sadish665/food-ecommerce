package main

import (
	"net/http"
	"github.com/sadish665/food-ecommerce/responses"
	"github.com/sadish665/food-ecommerce/routes"
)

func main() {
	
	print("hello all, welcome to my new api")
	newRouter := routes.InitAllRoutes()
	print("starting server @8080")
	err := http.ListenAndServe(":8080", newRouter)
	if err != nil{
		print("error starting error:", err.Error())
	}
}

func HandleInitialRoutes(w http.ResponseWriter, r *http.Request){
	data := map[string]interface{}{
		"userList": []string{"user1","user2"},
		"productList":[]string{"product1","product2"},
		"totalLength":100,
		"page":4,
	}
	// responses.ErrorResponds(w, http.StatusInternalServerError, "Something went wroong")
	responses.SuccessResponds(w,data, "Successfully fetched user list, product list, length and page")
}