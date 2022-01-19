package entity

type User struct {
	Id       int     `json:"id" form:"id"`
	Name     *string `json:"name" form:"name"`
	Email    *string `json:"email" form:"email"`
	Password *string `json:"password" form:"password"`
	Address  *string `json:"address" form:"address"`
}
