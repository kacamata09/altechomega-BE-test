package usecase_test

import (
	"altech-omega-api/domain"
	"altech-omega-api/usecase"
	mockdata "altech-omega-api/tests/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking AuthorRepository
type MockAuthorRepo struct {
	mock.Mock
}

func (m *MockAuthorRepo) GetAll() ([]domain.Author, error) {
	args := m.Called()
	return args.Get(0).([]domain.Author), args.Error(1)
}

func (m *MockAuthorRepo) GetByID(id string) (domain.Author, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Author), args.Error(1)
}

func (m *MockAuthorRepo) GetByIDWithBooks(id string) (domain.Author, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Author), args.Error(1)
}

func (m *MockAuthorRepo) Create(input *domain.Author) error {
	args := m.Called(input)
	return args.Error(0)
}

func (m *MockAuthorRepo) Update(id string, input *domain.Author) error {
	args := m.Called(id, input)
	return args.Error(0)
}

func (m *MockAuthorRepo) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestAuthorUsecase_GetAll(t *testing.T) {
	mockRepo := new(MockAuthorRepo)
	authors := mockdata.Authors 

	mockRepo.On("GetAll").Return(authors, nil)

	usecase := usecase.CreateAuthorUseCase(mockRepo)

	result, err := usecase.GetAll()
	assert.NoError(t, err)
	assert.Len(t, result, len(authors))
	assert.Equal(t, "Anshar", result[0].Name)
	assert.Equal(t, "Si Golang", result[1].Name)
	mockRepo.AssertExpectations(t)
}

func TestAuthorUsecase_GetByID(t *testing.T) {
	mockRepo := new(MockAuthorRepo)
	author := mockdata.Authors[0] 

	mockRepo.On("GetByID", "1156e695-8043-4c51-a9ae-d05e0eba9b2c").Return(author, nil)

	usecase := usecase.CreateAuthorUseCase(mockRepo)

	result, err := usecase.GetByID("1156e695-8043-4c51-a9ae-d05e0eba9b2c")
	assert.NoError(t, err)
	assert.Equal(t, "Anshar", result.Name)
	mockRepo.AssertExpectations(t)
}

func TestAuthorUsecase_GetByIDWithBooks(t *testing.T) {
	mockRepo := new(MockAuthorRepo)
	author := mockdata.Authors[0]
	author.Books = mockdata.Books  

	mockRepo.On("GetByIDWithBooks", "1156e695-8043-4c51-a9ae-d05e0eba9b2c").Return(author, nil)

	usecase := usecase.CreateAuthorUseCase(mockRepo)

	result, err := usecase.GetByIDWithBooks("1156e695-8043-4c51-a9ae-d05e0eba9b2c")

	assert.NoError(t, err)
	assert.Equal(t, "Anshar", result.Name)
	assert.Equal(t, "Hyouka", result.Books[1].Title)
	assert.Len(t, result.Books, len(mockdata.Books))  
}


func TestAuthorUsecase_Create(t *testing.T) {
	mockRepo := new(MockAuthorRepo)
	// author := &domain.Author{
	// 	Name:        "Anshar",
	// 	Bio:         "Orangnya baik sekali",
	// 	BirthDate:   "2000-08-03",
	// 	Nationality: "Indonesia",
	// }
	author := &mockdata.Authors[0]

	mockRepo.On("Create", author).Return(nil)

	usecase := usecase.CreateAuthorUseCase(mockRepo)

	err := usecase.Create(author)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAuthorUsecase_Update(t *testing.T) {
	mockRepo := new(MockAuthorRepo)
	author := &domain.Author{
		Name:        "Anshar Anshar",
		Bio:         "Dia pakai kacamata",
		BirthDate:   "1945-08-17",
		Nationality: "Jepang",
	}

	mockRepo.On("Update", "1156e695-8043-4c51-a9ae-d05e0eba9b2c", author).Return(nil)

	usecase := usecase.CreateAuthorUseCase(mockRepo)

	err := usecase.Update("1156e695-8043-4c51-a9ae-d05e0eba9b2c", author)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAuthorUsecase_Delete(t *testing.T) {
	mockRepo := new(MockAuthorRepo)

	mockRepo.On("Delete", "1156e695-8043-4c51-a9ae-d05e0eba9b2c").Return(nil)

	usecase := usecase.CreateAuthorUseCase(mockRepo)

	err := usecase.Delete("1156e695-8043-4c51-a9ae-d05e0eba9b2c")
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
