package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ErrorJson struct {
	StatusCode   int    `json:"status_code"`
	ErrorMessage string `json:"error_message"`
}

func (e *ErrorJson) Json() ([]byte, error) {
	data, err := json.Marshal(e)

	if err != nil {
		return nil, fmt.Errorf("failed to encode order: %w", err)
	}
	return data, nil
}

type Order struct {
	OrderId     uint64     `json:"order_id"`
	CustomerID  uuid.UUID  `json:"customer_id"`
	LineItems   []LineItem `json:"line_items"`
	OrderStatus string     `json:"order_status"`

	CreatedAt   time.Time `json:"created_at"`
	ShippedAt   time.Time `json:"shipped_at"`
	CompletedAt time.Time `json:"completed_at"`
}

type LineItem struct {
	ItemID   uuid.UUID `json:"item_id"`
	Quantity uint      `json:"quantity"`
	Price    float64   `json:"price"`
}

// repository
