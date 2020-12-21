package repository

import "github.com/vanilla/go-jwt-crud/api/entities"

type UserRepository interface {
	FindAll() ([]entities.User, error)
	FindById(uint64) (entities.User, error)
	Save(entities.User) (bool, error)
	Update(uint64, entities.User) (bool, error)
	Delete(uint64) (bool, error)
	Login(string) (entities.User, error)
}