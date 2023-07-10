package domains

type User struct {
	BaseEntity
	Name string `json:"name,omitempty" validate:"required"`
}
