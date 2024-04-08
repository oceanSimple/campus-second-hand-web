package model

type Account struct {
	Id        uint64 `json:"id" gorm:"id"`
	Account   string `json:"account" gorm:"account"`
	Password  string `json:"password" gorm:"password"`
	Phone     string `json:"phone" gorm:"phone"`
	Email     string `json:"email" gorm:"email"`
	Nickname  string `json:"nickname" gorm:"nickname"`
	Age       uint8  `json:"age" gorm:"age"`
	Gender    uint8  `json:"gender" gorm:"gender"`
	Campus    string `json:"campus" gorm:"campus"`
	Dormitory string `json:"dormitory" gorm:"dormitory"`
	Common
}

func (a *Account) TableName() string {
	return "account"
}

//func (a *Account) AfterFind() (err error) {
//	if a.IsDelete == 1 {
//		err = gorm.ErrRecordNotFound
//	}
//	return
//}
