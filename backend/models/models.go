package models

import (
	"time"
)

// Item represents a product in the store that is available for purchase.
type Item struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CreatedAt   time.Time `json:"created_at"`
}

// OrderItem is item that is currently in the order.
type OrderItem struct {
	ItemID   int     `json:"item_id"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

// Order represents the complete purchase submitted
type Order struct {
	ID     int         `json:"id"`
	UserID int         `json:"user_id"`
	Items  []OrderItem `json:"items"`
	Total  float64     `json:"total"`
	Status string      `json:"status"` // pending, paid, failed
}

// Cart is a single cart row representing one item in a user's cart.
type Cart struct {
	UserID   int `json:"user_id"`
	ItemID   int `json:"item_id"`
	Quantity int `json:"quantity"`
}

// CartItem represents a cart item with additional details for response.
type CartItem struct {
	ItemID      int       `json:"item_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`      // Price in cents
	Stock       int       `json:"stock"`      // Current stock level
	CreatedAt   time.Time `json:"created_at"` // Time when the item was added to the cart
	Quantity    int       `json:"quantity"`
}

// User represents a user account.
type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Hash      []byte    `json:"-"` // Password hash, never sent in responses
	CreatedAt time.Time `json:"created_at"`
}
