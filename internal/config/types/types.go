package types

type Student struct {
	Id    int64  `json:"id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
	Age   int    `json:"age" validate:"required"`
}

// For updated student
type UpdateStudent struct {
	Name  *string `json:"name" validate:"omitempty"`
	Email *string `json:"email" validate:"omitempty"`
	Age   *int    `json:"age" validate:"omitempty"`
}
