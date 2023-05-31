package models

type Book struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Title      string `json:"title"`
	AuthorName string `json:"author_name" gorm:"column: author_name"`
}
