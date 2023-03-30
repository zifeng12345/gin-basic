package tables

import "time"

type Waiting struct {
	ID          uint64     `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt   time.Time  `json:"createdAt" sql:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time  `sql:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt   *time.Time `gorm:"null" json:"deletedAt"`
	StoreId     string     `json:"storeId"`
	UserName    string     `json:"userName"`
	Desc        string     `json:"desc"`
	PhoneNumber string     `json:"phoneNumber"`
	Status      int32      `json:"status"`
	Number      int32      `json:"number"`
	Day         string     `json:"day"`
	Pax         int32      `json:"pax"`
}
