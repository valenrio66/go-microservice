package model

type Book struct {
	BooksID int    `gorm:"primaryKey;autoIncrement" json:"books_id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
}

// TableName overrides the table name used by Book to `book`
func (Book) TableName() string {
	return "book"
}
