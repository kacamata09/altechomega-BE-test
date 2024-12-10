package domain

type Book struct {
	ID              string         `json:"id"`
	Title       	string         `json:"title"`
	AuthorID		string         `json:"author_id"`
	PublishDate     string         `json:"publish_date"`
	Description     string         `json:"description"`
	Pages   		int16		   `json:"pages"`
	Genre       	string         `json:"genre"`
	CreatedAt       string         `json:"created_at"`
	UpdatedAt       string         `json:"updated_at"`
}



type BookRepository interface {
	GetAll() ([]Book, error)
	GetByID(id string) (Book, error)
	Create(book *Book) (error)
	Update(id string, book *Book) error
	Delete(id string) error
}

type BookUsecase interface {
	GetAll() ([]Book, error)
	GetByID(id string) (Book, error)
	Create(book *Book) error
	Update(id string, book *Book) error
	Delete(id string) error
}
