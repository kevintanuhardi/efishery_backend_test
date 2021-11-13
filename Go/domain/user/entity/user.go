package entity

import "time"

type Order struct {
	ID         int       `json:"id"`
	CustomerID int       `json:"customer_id"`
	TrxID      string    `json:"trx_id"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type OrderStore struct {
	ID         int       `json:"id"`
	OrderID    int       `json:"order_id"`
	StoreID    int       `json:"store_id"`
	StoreTrxID string    `json:"store_trx_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type OrderStoreItem struct {
	ID           int       `json:"id"`
	OrderStoreID int       `json:"order_store_id"`
	SKU          string    `json:"sku"`
	Name         string    `json:"name"`
	Uom          string    `json:"uom"`
	Quantity     float64   `json:"quantity"`
	PriceUnit    float64   `json:"price_unit"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type OrderAggregate struct {
	ID         int
	CustomerID int
	TrxID      string
	Status     string

	OrderStoreID int
	StoreID      int
	StoreTrxID   string

	OrderStoreItemID int
	SKU              string
	Name             string
	Uom              string
	Quantity         float64
	PriceUnit        float64

	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID		int					`json:"id"`
	Phone string			`json:"phone"`
	Name	string			`json:"name"`
	Role	string			`json:"role"`
	Password	string	`json:"password"`
}

func (User) TableName() string {
	return "user"
}
