package dto

type UserResponseModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserCreateRequestModel struct {
	Name string `json:"name" validate:"required"`
}
