package model

type Product struct {
	Id          uint64  `json:"id" gorm:"id"`
	Title       string  `json:"title" gorm:"title"`
	Description string  `json:"description" gorm:"description"`
	Price       float64 `json:"price" gorm:"price"`
	Stock       uint64  `json:"stock" gorm:"stock"`
	Campus      string  `json:"campus" gorm:"campus"`
	Address     string  `json:"address" gorm:"address"`
	SellerID    uint64  `json:"sellerId" gorm:"seller_id"`
	Picture1    string  `json:"picture1" gorm:"picture1"`
	Picture2    string  `json:"picture2" gorm:"picture2"`
	Picture3    string  `json:"picture3" gorm:"picture3"`
	Type1       uint64  `json:"type1" gorm:"type1"`
	Type2       uint64  `json:"type2" gorm:"type2"`
	Others      string  `json:"others" gorm:"others"`
	Common
}

// TableName returns the table name of the Product model.
func (Product) TableName() string {
	return "product"
}
