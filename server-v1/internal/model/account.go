package model

type Account struct {
	Id       uint64 `json:"id" gorm:"id"`
	Account  string `json:"account" gorm:"account"`
	Password string `json:"password" gorm:"password"`
	Phone    string `json:"phone" gorm:"phone"`
	Email    string `json:"email" gorm:"email"`
}

func (a *Account) TableName() string {
	return "account"
}
