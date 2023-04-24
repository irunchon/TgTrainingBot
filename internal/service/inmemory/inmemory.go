package inmemory

import (
	"errors"

	"github.com/irunchon/TgTrainingBot/internal/service/model"
)

var ErrNotFound = errors.New("not found")

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []model.Product {
	return model.AllProducts
}

func (s *Service) Get(ID uint64) (*model.Product, error) {
	if ID > uint64(len(model.AllProducts)-1) || ID < 0 {
		return nil, ErrNotFound
	}
	return &model.AllProducts[ID], nil
}

func (s *Service) Create(newProduct model.Product) uint64 {
	model.AllProducts = append(model.AllProducts, newProduct)
	return uint64(len(model.AllProducts) - 1)
}

func (s *Service) Remove(ID uint64) error {
	if ID > uint64(len(model.AllProducts)-1) || ID < 0 {
		return ErrNotFound
	}
	model.AllProducts = append(model.AllProducts[:ID], model.AllProducts[ID+1:]...)
	return nil
}

func (s *Service) Update(ID uint64, update model.Product) error {
	if ID > uint64(len(model.AllProducts)-1) || ID < 0 {
		return ErrNotFound
	}
	model.AllProducts[ID].Title = update.Title
	return nil
}
