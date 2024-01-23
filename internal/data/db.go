package data

import (
	"time"

	"github.com/uptrace/bun"
)

type Company struct {
	bun.BaseModel
	ID          int64 `bun:",pk,autoincrement"`
	CompanyName string
}

type Customer struct {
	bun.BaseModel
	ID          int64 `bun:",pk,autoincrement"`
	UserID      string
	Login       string
	Password    []byte
	Name        string
	CompanyID   int64
	CreditCards []string `bun:",array"`
}

type Delivery struct {
	bun.BaseModel
	ID                int64 `bun:",pk,autoincrement"`
	OrderItemID       int64
	DeliveredQuantity int
}

type OrderItem struct {
	bun.BaseModel
	ID           int64 `bun:",pk,autoincrement"`
	OrderID      int64
	PricePerUnit float64
	Quantity     int64
	Product      string
}

type Order struct {
	bun.BaseModel
	ID             int64 `bun:",pk,autoincrement"`
	CreatedAt      time.Time
	OrderName      string
	CustomerUserID string
}
