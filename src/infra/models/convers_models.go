package models

/*
 * Author      : Jody (jody.almaida@gmail)
 * Modifier    :
 * Domain      : game-currency
 */

type ConversionRates struct {
	ID     int64   `db:"id"`
	FromID int64   `db:"currency_id_from"`
	ToID   int64   `db:"currency_id_to"`
	Rate   float64 `db:"rate"`
}
