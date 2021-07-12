package main

import (
	"lotr/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/lotrapp/book", controllers.Book)

	http.ListenAndServe(":8080", nil)
}
