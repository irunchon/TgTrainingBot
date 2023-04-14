package inmemory

import (
	"TgTrainingBot/internal/service/model"
	"errors"
	"fmt"
)

type InMemoryService struct{}

func NewService() *InMemoryService {
	return &InMemoryService{}
}

func (s *InMemoryService) List() []model.Product {
	return model.AllProducts
}

func (s *InMemoryService) Get(idx int) (*model.Product, error) {
	if idx > len(model.AllProducts)-1 {
		return nil, errors.New(fmt.Sprintf("idx = %d is out of range", idx))
	}
	if idx < 0 {
		return nil, errors.New(fmt.Sprintf("idx = %d is negative", idx))
	}
	return &model.AllProducts[idx], nil
}
