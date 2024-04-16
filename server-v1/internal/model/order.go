package model

type HistoryOrder struct {
	Id        uint64  `json:"id" gorm:"id"`
	ProductId uint64  `json:"productId" gorm:"product_id"`
	BuyerId   uint64  `json:"buyerId" gorm:"buyer_id"`
	SellerId  uint64  `json:"sellerId" gorm:"seller_id"`
	Title     string  `json:"title" gorm:"title"`
	Price     float64 `json:"price" gorm:"price"`
	Common
}

func (h *HistoryOrder) TableName() string {
	return "history_order"
}
