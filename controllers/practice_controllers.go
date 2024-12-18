package controllers

import (
	"net/http"

	"github.com/sadish665/food-ecommerce/responses"
)

func HandleSuccessRequest(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"users" : []string{"user1","user2"},
		"products" : []string{"product1","product2"},
		"totalLength" : 100,
		"page" : 4,
	}
	responses.SuccessResponds(w, data, "Successfully fetched user list, product list, length and page")
}

func HandleErrorRequest(w http.ResponseWriter, r *http.Request){
	responses.ErrosResponds(w, http.StatusInternalServerError, "Something went wrong")
}