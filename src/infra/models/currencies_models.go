package models

/*
 * Author      : Jody (jody.almaida@gmail)
 * Modifier    :
 * Domain      : game-currency
 */

type Currencies struct {
	ID   int64  `gorm:"id"`
	Name string `gorm:"name"`
}
