package usecase

import (
	"database/sql"
	"altech-omega-api/domain"
)

type BookUsecase struct {
	BookRepo domain.BookRepository
	DB       *sql.DB
}

func CreateBookUseCase(repo domain.BookRepository) domain.BookUsecase {
	usecase := BookUsecase{
		BookRepo: repo,
	}

	return &usecase
}

func (uc BookUsecase) GetAll() ([]domain.Book, error) {
	data, err := uc.BookRepo.GetAll()
	return data, err
}

func (uc BookUsecase) GetByID(id string) (domain.Book, error) {
	data, err := uc.BookRepo.GetByID(id)
	return data, err
}

func (uc BookUsecase) Create(input *domain.Book) error {
	err := uc.BookRepo.Create(input)
	return err
}

func (uc BookUsecase) Update(id string, input *domain.Book) error {
	err := uc.BookRepo.Update(id, input)
	return err
}

func (uc BookUsecase) Delete(id string) error {
	err := uc.BookRepo.Delete(id)
	return err
}
