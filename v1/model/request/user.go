package request

type UserLogin struct {
	Password string `json:"password" form:"password" binding:"required"`
	Email    string `json:"email"`
}
