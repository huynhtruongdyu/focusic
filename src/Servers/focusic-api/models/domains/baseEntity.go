package domains

import "time"

type BaseEntity struct {
	CreatedAt time.Time `json:"createdAt,omitempty"`
	CreatedBy string    `json:"createdBy,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	UpdatedBy string    `json:"updatedBy,omitempty"`
	IsDeleted bool      `json:"isDeleted,omitempty"`
}
