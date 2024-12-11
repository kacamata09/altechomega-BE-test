package repositoryMySql_test

import (
	"altech-omega-api/domain"
	repositoryMySql "altech-omega-api/repository/mysql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestBookRepo_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "author_id", "publish_date", "description", "pages", "genre", "created_at", "updated_at"}).
		AddRow("1156e695-8043-4c51-a9ae-d05e0eba9b2c", "The 5AM Club", "1156e695-8043-4c51-a9ae-d05e0eba9b2c", "2010-10-01", "Bangun pagi supaya sukses", 100, "Self Development", "2024-12-10", "2024-12-10").
		AddRow("35a16b5d-8eca-4055-a4d3-a62a355aeab8", "Hyouka", "1156e695-8043-4c51-a9ae-d05e0eba9b2c", "2005-08-04", "Oreki Houtarou", 150, "Light Novel", "2024-12-10", "2024-12-10")
	mock.ExpectQuery("SELECT \\* FROM books").WillReturnRows(rows)

	repo := repositoryMySql.CreateBookRepo(db)

	books, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, books, 2)
	assert.Equal(t, "The 5AM Club", books[0].Title)
	assert.Equal(t, "Hyouka", books[1].Title)
}

func TestBookRepo_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	row := sqlmock.NewRows([]string{"id", "title", "author_id", "publish_date", "description", "pages", "genre", "created_at", "updated_at"}).
		AddRow("1156e695-8043-4c51-a9ae-d05e0eba9b2c", "The 5AM Club", "1156e695-8043-4c51-a9ae-d05e0eba9b2c", "2010-10-01", "Bangun pagi supaya sukses", 100, "Self Development", "2024-12-10", "2024-12-10")
	mock.ExpectQuery("SELECT \\* FROM books where id=\\?").WithArgs("1156e695-8043-4c51-a9ae-d05e0eba9b2c").WillReturnRows(row)

	repo := repositoryMySql.CreateBookRepo(db)

	book, err := repo.GetByID("1156e695-8043-4c51-a9ae-d05e0eba9b2c")
	assert.NoError(t, err)
	assert.Equal(t, "The 5AM Club", book.Title)
	assert.Equal(t, "Self Development", book.Genre)
}

func TestBookRepo_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec("INSERT INTO books").
		WithArgs(sqlmock.AnyArg(), "The 5AM Club", "1156e695-8043-4c51-a9ae-d05e0eba9b2c", "2010-10-01", "Bangun pagi supaya sukses", 100, "Self Development").
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repositoryMySql.CreateBookRepo(db)

	book := &domain.Book{
		Title:       "The 5AM Club",
		AuthorID:    "1156e695-8043-4c51-a9ae-d05e0eba9b2c",
		PublishDate: "2010-10-01",
		Description: "Bangun pagi supaya sukses",
		Pages:       100,
		Genre:       "Self Development",
	}

	err = repo.Create(book)
	assert.NoError(t, err)
}

func TestBookRepo_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec("UPDATE books SET title = \\?, author_id = \\?, publish_date = \\?, description = \\?, pages = \\?, genre = \\? WHERE id = \\?").
		WithArgs("The 5AM Club Updated", "1156e695-8043-4c51-a9ae-d05e0eba9b2c", "2011-10-01", "Bangun pagi agar lebih sukses", 120, "Self Development", "1156e695-8043-4c51-a9ae-d05e0eba9b2c").
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := repositoryMySql.CreateBookRepo(db)

	book := &domain.Book{
		Title:       "The 5AM Club Updated",
		AuthorID:    "1156e695-8043-4c51-a9ae-d05e0eba9b2c",
		PublishDate: "2011-10-01",
		Description: "Bangun pagi agar lebih sukses",
		Pages:       120,
		Genre:       "Self Development",
	}

	err = repo.Update("1156e695-8043-4c51-a9ae-d05e0eba9b2c", book)
	assert.NoError(t, err)
}

func TestBookRepo_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec("DELETE FROM books WHERE id = \\?").
		WithArgs("1156e695-8043-4c51-a9ae-d05e0eba9b2c").
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := repositoryMySql.CreateBookRepo(db)

	err = repo.Delete("1156e695-8043-4c51-a9ae-d05e0eba9b2c")
	assert.NoError(t, err)
}
