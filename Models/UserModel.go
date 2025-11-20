package Models

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegister struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" `
	Address  string `json:"address"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginRes struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone" `
	Address  string `json:"address"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func (b *User) TableName() string {
	return "User"
}
