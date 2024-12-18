package controllers

import (
	"net/http"

	"github.com/sadish665/food-ecommerce/responses"
)

func HandleGetAllUsers(w http.ResponseWriter, r *http.Request){
	users := []string{"user1","user2"}
	responses.SuccessResponds(w, users, "Successfully fetched user list")
}