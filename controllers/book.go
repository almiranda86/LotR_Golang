package controllers

import (
	"encoding/json"
	"fmt"
	"lotr/service"
	"net/http"
)

func Book(w http.ResponseWriter, r *http.Request) {
	response := service.ListAllBooks()

	fmt.Println(response)
	json.NewEncoder(w).Encode(response)
}
