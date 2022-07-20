package model

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ProductID uuid.UUID `json:"product_id"`
	Name      string    `json:"name"`
	Quantity  int64     `json:"quantity"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type ProductRequest struct {
	ProductID uuid.UUID `json:"product_id"`
	Name      string    `json:"name"`
	Quantity  int64     `json:"quantity"`
}

type ProductResponse struct {
	Status   string   `json:"status"`
	Response *Product `json:"response,omitempty"`
}
