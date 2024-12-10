package repositoryMySql

import (
	"database/sql"
	"fmt"
	"altech-omega-api/domain"
	"github.com/google/uuid"
)

type bookRepo struct {
	DB *sql.DB
}

func CreateBookRepo(db *sql.DB) domain.BookRepository {
	return &bookRepo{DB: db}
}

func (repo *bookRepo) GetAll() ([]domain.Book, error) {

	rows, err := repo.DB.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []domain.Book

	for rows.Next() {
		var book domain.Book
		err := rows.Scan(&book.ID, &book.Title, &book.AuthorID, &book.PublishDate,
			&book.Description, &book.Pages, &book.Genre, &book.CreatedAt, &book.UpdatedAt)

		if err != nil {
			return data, err
		}
		data = append(data, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return data, err
}

func (repo *bookRepo) GetByID(id string) (domain.Book, error) {
	row := repo.DB.QueryRow("SELECT * FROM books where id=?", id)
	fmt.Println(id)

	var data domain.Book

	err := row.Scan(&data.ID, &data.Title, &data.AuthorID, &data.PublishDate,
		&data.Description, &data.Pages, &data.Genre, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		return data, err
	}

	if err := row.Err(); err != nil {
		return data, err
	}
	// fmt.Println(data)
	return data, err
}


func (repo *bookRepo) Create(book *domain.Book) error {
	newUUID, _ := uuid.NewRandom()
	// newUUID, _ := uuid.NewUUID()
	id := newUUID.String()

	_, err := repo.DB.Exec("INSERT INTO books (id, title, author_id, publish_date, description, pages, genre) values (?, ?, ?, ?, ?, ?, ?)",
		id, book.Title, book.AuthorID, book.PublishDate,
		book.Description, book.Pages, book.Genre)
	return err
}

func (repo *bookRepo) Update(id string, book *domain.Book) error {

	_, err := repo.DB.Exec("UPDATE books SET title = ?, author_id = ?, publish_date = ?, description = ?, pages = ?, genre = ? WHERE id = ?",
		book.Title, book.AuthorID, book.PublishDate,
		book.Description, book.Pages, book.Genre, id)
	return err
}

func (repo *bookRepo) Delete(id string) error {

	_, err := repo.DB.Exec("DELETE FROM books WHERE id = ?", id)

	return err
}
