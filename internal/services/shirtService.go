package services

import (
	"github.com/yosa12978/MyShirts/internal/models"
	"github.com/yosa12978/MyShirts/internal/repos"
)

type ShirtService interface {
	GetShirts() ([]models.Shirt, error)
	GetShirtByID(id string) (models.Shirt, error)
	AddShirt(shirt models.Shirt) error
	UpdateShirt(shirt models.Shirt) error
	DeleteShirt(id string) error
}

type shirtService struct {
	shirtRepository repos.ShirtRepo
}

func NewShirtService(repo repos.ShirtRepo) ShirtService {
	return &shirtService{
		shirtRepository: repo,
	}
}

func (ss *shirtService) GetShirts() ([]models.Shirt, error) {
	return ss.shirtRepository.GetAll()
}

func (ss *shirtService) GetShirtByID(id string) (models.Shirt, error) {
	return ss.shirtRepository.GetByID(id)
}

func (ss *shirtService) AddShirt(shirt models.Shirt) error {
	return ss.shirtRepository.Create(shirt)
}

func (ss *shirtService) UpdateShirt(shirt models.Shirt) error {
	return ss.shirtRepository.Update(shirt)
}

func (ss *shirtService) DeleteShirt(id string) error {
	return ss.shirtRepository.Delete(id)
}
