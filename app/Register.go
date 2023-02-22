package app

type Register struct {
	Username string `json:"username" form:"username" binding:"required" validate:"min:1"`
	Email    string `json:"email" form:"email" binding:"required" validate:"email"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:6"`
}
