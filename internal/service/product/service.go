package product

import (
	"errors"
	"fmt"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}

func (s *Service) Get(idx int) (*Product, error) {
	if idx > len(allProducts)-1 || idx < 0 {
		return nil, errors.New(fmt.Sprintf("Wrong product idx = %d", idx))
	}
	return &allProducts[idx], nil
}
