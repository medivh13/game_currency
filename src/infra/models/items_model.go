package models

/*
 * Author      : Jody (jody.almaida@gmail)
 * Modifier    :
 * Domain      : game-currency
 */

import "time"

type Items struct {
	ID        int64     `gorm:"id"`
	SKU       string    `gorm:"sku"`
	Name      string    `gorm:"name"`
	Price     float64   `gorm:"price"`
	Qty       float64   `gorm:"inventory_qty"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}
