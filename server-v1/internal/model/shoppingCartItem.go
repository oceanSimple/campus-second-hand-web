package model

import "time"

type ShoppingCartItem struct {
	ProductId uint64    `json:"productId"` // 商品ID
	Picture   string    `json:"picture"`   // 商品图片
	Title     string    `json:"title"`     // 商品标题
	Stock     uint64    `json:"stock"`     // 库存
	Price     float64   `json:"price"`     // 单价
	Quantity  uint64    `json:"quantity"`  // 数量
	Total     float64   `json:"total"`     // 总金额
	Time      time.Time `json:"time"`      // 添加时间
}
