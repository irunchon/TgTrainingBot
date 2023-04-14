package service

import "TgTrainingBot/internal/service/model"

type Service interface {
	List() []model.Product
	Get(ID uint64) (*model.Product, error)
	Create(model.Product) uint64
	Remove(ID uint64) (bool, error)
}
