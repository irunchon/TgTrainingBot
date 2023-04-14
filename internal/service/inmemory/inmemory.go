package inmemory

import (
	"TgTrainingBot/internal/service/model"
	"errors"
)

var ErrNotFound = errors.New("not found")

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []model.Product {
	return model.AllProducts
}

func (s *Service) Get(idx uint64) (*model.Product, error) {
	if idx > uint64(len(model.AllProducts)-1) || idx < 0 {
		return nil, ErrNotFound
	}
	return &model.AllProducts[idx], nil
}
