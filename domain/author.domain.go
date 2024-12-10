package domain

type Author struct {
	ID              string         `json:"id"`
	Name       		string         `json:"name"`
	Bio				string         `json:"bio"`
	BirthDate       string         `json:"birth_date"`
	Nationality     string         `json:"nationality"`
	Books   		[]Book		   `json:"book"`
	CreatedAt       string         `json:"created_at"`
	UpdatedAt       string         `json:"updated_at"`
}



type AuthorRepository interface {
	GetAll() ([]Author, error)
	GetByID(id string) (Author, error)
	GetByIDWithBooks(id string) (Author, error)
	Create(author *Author) (error)
	Update(id string, author *Author) error
	Delete(id string) error
}

type AuthorUsecase interface {
	GetAll() ([]Author, error)
	GetByID(id string) (Author, error)
	GetByIDWithBooks(id string) (Author, error)
	Create(author *Author) error
	Update(id string, author *Author) error
	Delete(id string) error
}
