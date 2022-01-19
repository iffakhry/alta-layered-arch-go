package user

import (
	_entity "sirclo/restapi/layered/entity"
)

type UserRequestFormat struct {
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Address string `json:"address" form:"address"`
}

type UserResponseFormat struct {
	Id      int     `json:"id" form:"id"`
	Name    *string `json:"name" form:"name"`
	Email   *string `json:"email" form:"email"`
	Address *string `json:"address" form:"address"`
}

func (req *UserRequestFormat) ToEntity() *_entity.User {
	return &_entity.User{
		Name:    &req.Name,
		Email:   &req.Email,
		Address: &req.Address,
	}
}

func FromEntity(entity _entity.User) UserResponseFormat {
	return UserResponseFormat{
		Id:      entity.Id,
		Name:    entity.Name,
		Email:   entity.Email,
		Address: entity.Address,
		// UpdatedAt: domain.UpdatedAt,
		// DeletedAt: domain.DeletedAt,
	}
}
