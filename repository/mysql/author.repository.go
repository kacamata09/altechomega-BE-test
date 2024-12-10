package repositoryMySql

import (
	"altech-omega-api/domain"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type authorRepo struct {
	DB *sql.DB
}

func CreateAuthorRepo(db *sql.DB) domain.AuthorRepository {
	return &authorRepo{DB: db}
}

func (repo *authorRepo) GetAll() ([]domain.Author, error) {

	rows, err := repo.DB.Query("SELECT * FROM authors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []domain.Author

	for rows.Next() {
		var author domain.Author
		err := rows.Scan(&author.ID, &author.Name, &author.Bio, &author.BirthDate,
			&author.Nationality, &author.CreatedAt, &author.UpdatedAt)

		if err != nil {
			return data, err
		}
		data = append(data, author)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return data, err
}

func (repo *authorRepo) GetByID(id string) (domain.Author, error) {
	row := repo.DB.QueryRow("SELECT * FROM authors where id=?", id)
	fmt.Println(id)

	var data domain.Author

	err := row.Scan(&data.ID, &data.Name, &data.Bio, &data.BirthDate,
		&data.Nationality, &data.CreatedAt, &data.UpdatedAt)
	if err != nil {
		return data, err
	}

	if err := row.Err(); err != nil {
		return data, err
	}
	// fmt.Println(data)
	return data, err
}

func (repo *authorRepo) GetByIDWithBooks(id string) (domain.Author, error) {
	query := `
	SELECT 
		au.id,
		au.name,
		au.bio,
		au.nationality,
		b.id,
		b.title,
		b.author_id,
		b.publish_date,
		b.description,
		b.pages,
		b.genre
	 FROM 
	 	authors au
	 LEFT JOIN
	 	books b
	 ON
	 	b.author_id = au.id
	 WHERE id=?
	`
	rows, err := repo.DB.Query(query, id)
	fmt.Println(id)

	var data domain.Author
	var books []domain.Book

	for rows.Next() {
		var book domain.Book

		err := rows.Scan(&data.ID, &data.Name, &data.Bio, &data.BirthDate,
			&data.Nationality, &book.ID, &book.Title, &book.AuthorID, &book.PublishDate,
			&book.Description, &book.Pages, &book.Genre)
		if err != nil {
			return data, err
		}

		books = append(books, book)
	}

	data.Books = books

	if err := rows.Err(); err != nil {
		return data, err
	}
	// fmt.Println(data)
	return data, err
}

func (repo *authorRepo) Create(author *domain.Author) error {
	newUUID, _ := uuid.NewRandom()
	// newUUID, _ := uuid.NewUUID()
	id := newUUID.String()

	_, err := repo.DB.Exec("INSERT INTO authors (id, name, bio, birth_date, nationality) values (?, ?, ?, ?, ?)",
		id, author.Name, author.Bio, author.BirthDate, author.Nationality)
	return err
}

func (repo *authorRepo) Update(id string, author *domain.Author) error {

	_, err := repo.DB.Exec("UPDATE authors SET name = ?, bio = ?, birth_date = ?, nationality = ? WHERE id = ?",
		author.Name, author.Bio, author.BirthDate, author.Nationality, id)
	return err
}

func (repo *authorRepo) Delete(id string) error {

	_, err := repo.DB.Exec("DELETE FROM authors WHERE id = ?", id)

	return err
}
