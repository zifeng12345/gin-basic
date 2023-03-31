package tables

import "time"

type Bases struct {
	ID        uint64     `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time  `json:"createdAt" sql:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `sql:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"null" json:"deletedAt"`
}
