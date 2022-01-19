package user

import "sirclo/restapi/layered/entity"

type User interface {
	GetAll() ([]entity.User, error)
	Get(int) (entity.User, error)
	Create(entity.User) (int, error)
	Update(entity.User) (int, error)
	Delete(int) (int, error)
}
