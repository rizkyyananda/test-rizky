package dto

type LoginDTO struct {
	UserName string `json:"userName" form:"userName" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
