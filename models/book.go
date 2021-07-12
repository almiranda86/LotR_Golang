package models

type Docs struct {
	Docs []Book
}

type Book struct {
	Id   string `json:"_id"`
	Name string `json:"name"`
}

// constructor function
func (b *Book) Constructor() {
	b.Id = ""
	b.Name = ""
}
