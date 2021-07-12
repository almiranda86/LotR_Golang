package service

import (
	"encoding/json"
	"lotr/data"
	"lotr/models"
)

func ListAllBooks() []models.Book {
	docs := models.Docs{}

	responseQuery := data.ListAllBooks()

	json.Unmarshal([]byte(responseQuery), &docs)

	return docs.Docs
}
