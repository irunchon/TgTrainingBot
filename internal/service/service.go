package service

import "TgTrainingBot/internal/service/model"

type Service interface {
	List() []model.Product
	Get(ID uint64) (*model.Product, error)
}

/*
	Describe(ID uint64) (*{domain}., error) // get?
	List(cursor uint64, limit uint64) ([]{domain}., error)
	Create({domain}.) (uint64, error) // new
	Update(ID uint64,  {domain}.) error // edit
	Remove(ID uint64) (bool, error) // delete
*/
