package models

// Base model for all the models
type Base struct {
	ID        int64  `json:"id" orm:"pk;auto;column(id)"`
	CreatedAt *int64 `json:"created_at"`
	UpdatedAt *int64 `json:"updated_at"`
	DeletedAt *int64 `json:"deleted_at"`
}
