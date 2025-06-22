package domain

import "time"

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	NameProduct string `json:"nameProduct" binding:"required"`
	Desciption  string
	Price       float64
	Stock       int
	UpdatedAt   time.Time
	CreatedAt   time.Time
}
