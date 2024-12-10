package usecase

import (
	"database/sql"
	"altech-omega-api/domain"
)

type AuthorUsecase struct {
	AuthorRepo domain.AuthorRepository
	DB         *sql.DB
}

func CreateAuthorUseCase(repo domain.AuthorRepository) domain.AuthorUsecase {
	usecase := AuthorUsecase{
		AuthorRepo: repo,
	}

	return &usecase
}

func (uc AuthorUsecase) GetAll() ([]domain.Author, error) {
	data, err := uc.AuthorRepo.GetAll()
	return data, err
}

func (uc AuthorUsecase) GetByID(id string) (domain.Author, error) {
	data, err := uc.AuthorRepo.GetByID(id)
	return data, err
}

func (uc AuthorUsecase) GetByIDWithBooks(id string) (domain.Author, error) {
	data, err := uc.AuthorRepo.GetByIDWithBooks(id)
	return data, err
}

func (uc AuthorUsecase) Create(input *domain.Author) error {
	err := uc.AuthorRepo.Create(input)
	return err
}

func (uc AuthorUsecase) Update(id string, input *domain.Author) error {
	err := uc.AuthorRepo.Update(id, input)
	return err
}

func (uc AuthorUsecase) Delete(id string) error {
	err := uc.AuthorRepo.Delete(id)
	return err
}
