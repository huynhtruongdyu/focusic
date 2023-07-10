package domains

import "time"

type BaseEntity struct {
	ID uint64 `json:"id" gorm:"primary_key;auto_increment;not_null"`
	//CreatedAt time.Time `json:"createdAt,omitempty" gorm:"default:CURRENT_TIMESTAMP()"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	CreatedBy string    `json:"createdBy,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	UpdatedBy string    `json:"updatedBy,omitempty"`
	IsDeleted bool      `json:"isDeleted,omitempty"`
}
