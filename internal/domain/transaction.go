package domain

import "time"

type Transaction struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	ProductId  uint
	Quantity   int
	TotalPrice float64
	Date       time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time
	CreatedAt  time.Time
}
