package repositoryMySql_test

import (
	"altech-omega-api/domain"
	repositoryMySql "altech-omega-api/repository/mysql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestAuthorRepo_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Mock query and result
	rows := sqlmock.NewRows([]string{"id", "name", "bio", "birth_date", "nationality", "created_at", "updated_at"}).
		AddRow("1156e695-8043-4c51-a9ae-d05e0eba9b2c", "Anshar", "Orangnya baik sekali", "2000-08-03", "Indonesia", "2024-12-10", "2024-12-10").
		AddRow("35a16b5d-8eca-4055-a4d3-a62a355aeab8", "Si Golang", "Anaknya pintar", "2007-03-05", "Amerika", "2024-12-10", "2024-12-10")
	mock.ExpectQuery("SELECT \\* FROM authors").WillReturnRows(rows)

	repo := repositoryMySql.CreateAuthorRepo(db)

	authors, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, authors, 2)
	assert.Equal(t, "Anshar", authors[0].Name)
	assert.Equal(t, "Si Golang", authors[1].Name)
}

func TestAuthorRepo_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	row := sqlmock.NewRows([]string{"id", "name", "bio", "birth_date", "nationality", "created_at", "updated_at"}).
		AddRow("1156e695-8043-4c51-a9ae-d05e0eba9b2c", "Anshar", "Orangnya baik sekali", "2000-08-03", "Indonesia", "2024-12-10", "2024-12-10")
	mock.ExpectQuery("SELECT \\* FROM authors where id=\\?").WithArgs("1").WillReturnRows(row)

	repo := repositoryMySql.CreateAuthorRepo(db)

	author, err := repo.GetByID("1")
	assert.NoError(t, err)
	assert.Equal(t, "Anshar", author.Name)
	assert.Equal(t, "Indonesia", author.Nationality)
}

// func TestAuthorRepo_GetByIDWithBooks(t *testing.T) {
// 	db, mock, err := sqlmock.New()
// 	assert.NoError(t, err)
// 	defer db.Close()

// 	row := sqlmock.NewRows([]string{"id", "name", "bio", "birth_date", "nationality", "created_at", "updated_at"}).
// 		AddRow("1156e695-8043-4c51-a9ae-d05e0eba9b2c", "Anshar", "Orangnya baik sekali", "2000-08-03", "Indonesia", "2024-12-10", "2024-12-10")
// 	mock.ExpectQuery("SELECT \\* FROM authors where id=\\?").WithArgs("1").WillReturnRows(row)

// 	repo := repositoryMySql.CreateAuthorRepo(db)

// 	author, err := repo.GetByIDWithBooks("1")
// 	assert.NoError(t, err)
// 	assert.Equal(t, "Anshar", author.Name)
// 	assert.Equal(t, "Indonesia", author.Nationality)
// }

func TestAuthorRepo_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec("INSERT INTO authors").
		WithArgs(sqlmock.AnyArg(), "Anshar", "Orangnya baik sekali", "2000-08-03", "Indonesia").
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repositoryMySql.CreateAuthorRepo(db)

	author := &domain.Author{
		Name:        "Anshar",
		Bio:         "Orangnya baik sekali",
		BirthDate:   "2000-08-03",
		Nationality: "Indonesia",
	}

	err = repo.Create(author)
	assert.NoError(t, err)
}

func TestAuthorRepo_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec("UPDATE authors SET name = \\?, bio = \\?, birth_date = \\?, nationality = \\? WHERE id = \\?").
		WithArgs("Updated Author", "Updated Bio", "2000-01-01", "Updated Nationality", "1156e695-8043-4c51-a9ae-d05e0eba9b2c").
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := repositoryMySql.CreateAuthorRepo(db)

	author := &domain.Author{
		Name:        "Updated Author",
		Bio:         "Updated Bio",
		BirthDate:   "2000-01-01",
		Nationality: "Updated Nationality",
	}

	err = repo.Update("1156e695-8043-4c51-a9ae-d05e0eba9b2c", author)
	assert.NoError(t, err)
}

func TestAuthorRepo_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec("DELETE FROM authors WHERE id = \\?").
		WithArgs("1156e695-8043-4c51-a9ae-d05e0eba9b2c").
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := repositoryMySql.CreateAuthorRepo(db)

	err = repo.Delete("1156e695-8043-4c51-a9ae-d05e0eba9b2c")
	assert.NoError(t, err)
}
