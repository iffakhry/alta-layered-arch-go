package user

import (
	"database/sql"
	_entity "sirclo/restapi/layered/entity"
)

type UserRequestFormat struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
}

type UserResponseFormat struct {
	Id      int    `json:"id" form:"id"`
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Address string `json:"address" form:"address"`
}

func (req *UserRequestFormat) ToEntity() *_entity.User {
	return &_entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Address:  sql.NullString{String: req.Address},
	}
}

func FromEntity(entity _entity.User) UserResponseFormat {
	return UserResponseFormat{
		Id:      entity.Id,
		Name:    entity.Name,
		Email:   entity.Email,
		Address: entity.Address.String,
		// UpdatedAt: domain.UpdatedAt,
		// DeletedAt: domain.DeletedAt,
	}
}
