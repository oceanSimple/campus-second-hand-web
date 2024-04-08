package model

import "time"

type Common struct {
	IsDeleted   uint8     `json:"isDelete" gorm:"is_deleted"`
	GmtCreate   time.Time `json:"gmtCreate" gorm:"gmt_create"`
	GmtModified time.Time `json:"gmtModified" gorm:"gmt_modified"`
}
