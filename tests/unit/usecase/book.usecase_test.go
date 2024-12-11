package usecase_test

import (
	"altech-omega-api/domain"
	"altech-omega-api/usecase"
	mockdata "altech-omega-api/tests/mock_data"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


type MockBookRepo struct {
	mock.Mock
}

func (m *MockBookRepo) GetAll() ([]domain.Book, error) {
	args := m.Called()
	return args.Get(0).([]domain.Book), args.Error(1)
}

func (m *MockBookRepo) GetByID(id string) (domain.Book, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Book), args.Error(1)
}

func (m *MockBookRepo) Create(input *domain.Book) error {
	args := m.Called(input)
	return args.Error(0)
}

func (m *MockBookRepo) Update(id string, input *domain.Book) error {
	args := m.Called(id, input)
	return args.Error(0)
}

func (m *MockBookRepo) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestBookUsecase_GetAll(t *testing.T) {
	mockRepo := new(MockBookRepo)
	books := mockdata.Books 

	mockRepo.On("GetAll").Return(books, nil)

	usecase := usecase.CreateBookUseCase(mockRepo)

	result, err := usecase.GetAll()
	assert.NoError(t, err)
	assert.Len(t, result, len(books))
	assert.Equal(t, "The 5AM Club", result[0].Title)
	assert.Equal(t, "Hyouka", result[1].Title)
	assert.Equal(t, "2dd2784b-31d5-4bb7-9429-644ba619476a", result[0].AuthorID) 
	assert.Equal(t, "b409d8d6-701a-4a31-b0b0-8db88f09f218", result[1].AuthorID) 
	mockRepo.AssertExpectations(t)
}


func TestBookUsecase_GetByID(t *testing.T) {
	mockRepo := new(MockBookRepo)
	book := mockdata.Books[0] 

	mockRepo.On("GetByID", "1156e695-8043-4c51-a9ae-d05e0eba9b2c").Return(book, nil)

	usecase := usecase.CreateBookUseCase(mockRepo)

	result, err := usecase.GetByID("1156e695-8043-4c51-a9ae-d05e0eba9b2c")
	assert.NoError(t, err)
	assert.Equal(t, "The 5AM Club", result.Title)
	assert.Equal(t, "2dd2784b-31d5-4bb7-9429-644ba619476a", result.AuthorID) 
	mockRepo.AssertExpectations(t)
}

func TestBookUsecase_Create(t *testing.T) {
	mockRepo := new(MockBookRepo)
	book := &mockdata.Books[0] 

	mockRepo.On("Create", book).Return(nil)

	usecase := usecase.CreateBookUseCase(mockRepo)

	err := usecase.Create(book)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}


func TestBookUsecase_Update(t *testing.T) {
	mockRepo := new(MockBookRepo)
	book := &domain.Book{
		Title:       "Hyouka - Misteri Dibalik Pembunuhan Ruang Tertutup",
		AuthorID:    "7a33d408-f9a2-4553-8791-1a6332a42ba4", 
		Description: "Oreki saat itu sedang disibukkan masalah dari club teater",
	}

	mockRepo.On("Update", "1156e695-8043-4c51-a9ae-d05e0eba9b2c", book).Return(nil)

	usecase := usecase.CreateBookUseCase(mockRepo)

	err := usecase.Update("1156e695-8043-4c51-a9ae-d05e0eba9b2c", book)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}


func TestBookUsecase_Delete(t *testing.T) {
	mockRepo := new(MockBookRepo)

	mockRepo.On("Delete", "1156e695-8043-4c51-a9ae-d05e0eba9b2c").Return(nil)

	usecase := usecase.CreateBookUseCase(mockRepo)

	err := usecase.Delete("1156e695-8043-4c51-a9ae-d05e0eba9b2c")
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
