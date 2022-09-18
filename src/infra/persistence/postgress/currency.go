package postgres

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : game-currency
 */

import (
	"errors"
	"log"

	dtoConversion "game-currency/src/app/dtos/conversion"
	dto "game-currency/src/app/dtos/currency"
	repositories "game-currency/src/domain/repositories"
	models "game-currency/src/infra/models"

	"github.com/jmoiron/sqlx"
)

const (
	GetConversionRate = `SELECT 
	id, 
	currency_id_from,
	currency_id_to,
	rate
	FROM game_currency.conversion_rates
	WHERE
	currency_id_from = $1 AND currency_id_to = $2
	OR
	currency_id_from = $3 AND currency_id_to = $4`

	GetCurrencyList = `SELECT 
	id, 
	name
	FROM game_currency.currencies
	WHERE
	deleted_at IS NULL
	`

	CreateCurrency = `INSERT INTO game_currency.currencies
	(name, created_at)
	VALUES 
	($1, now())
	`

	CreateConversionRate = `INSERT INTO game_currency.conversion_rates
	(currency_id_from, currency_id_to, rate, created_at)
	VALUES
	($1, $2, $3, now())
	`
)

var statement PreparedStatement

type PreparedStatement struct {
	getConversionResult  *sqlx.Stmt
	getCurrenciesList    *sqlx.Stmt
	createCurrency       *sqlx.Stmt
	createConversionRate *sqlx.Stmt
}

type ConversRepository struct {
	Connection *sqlx.DB
}

func NewConversionRepository(db *sqlx.DB) repositories.ConversionRateRepository {
	repo := &ConversRepository{db}
	InitPreparedStatement(repo)
	return repo
}

func (p *ConversRepository) Preparex(query string) *sqlx.Stmt {
	statement, err := p.Connection.Preparex(query)
	if err != nil {
		log.Fatalf("Failed to preparex query: %s. Error: %s", query, err.Error())
	}

	return statement
}

func InitPreparedStatement(m *ConversRepository) {
	statement = PreparedStatement{
		getConversionResult:  m.Preparex(GetConversionRate),
		getCurrenciesList:    m.Preparex(GetCurrencyList),
		createCurrency:       m.Preparex(CreateCurrency),
		createConversionRate: m.Preparex(CreateConversionRate),
	}
}

func (repo *ConversRepository) GetConversionRate(fromID, toID int64) (*models.ConversionRates, error) {
	var data []models.ConversionRates

	err := statement.getConversionResult.Select(&data, fromID, toID, toID, fromID)

	if err != nil {
		log.Println("Failed Query GetConversionRate : ", err.Error())
		return nil, err
	}

	if len(data) < 1 {
		err := errors.New("Data Not Found")
		log.Println("Failed Query GetConversionRate : ", err.Error())
		return nil, err
	}
	return &data[0], nil
}

func (repo *ConversRepository) GetCurrenciesList() (*[]models.Currencies, error) {
	var data []models.Currencies

	err := statement.getCurrenciesList.Select(&data)

	if err != nil {
		log.Println("Failed Query GetCurrencyList : ", err.Error())
		return nil, err
	}

	return &data, nil
}

func (p *ConversRepository) CreateCurrency(data *dto.CurrencyReqDTO) error {

	_, err := statement.createCurrency.Exec(data.Name)

	if err != nil {
		log.Println("Failed Query CreateCurrency : ", err.Error())
		return err
	}

	return nil
}

func (p *ConversRepository) CreateConversionRate(data *dtoConversion.CreateConversReqDTO) error {

	_, err := statement.createConversionRate.Exec(data.CurrencyIdFrom, data.CurrencyIdTo, data.Rate)

	if err != nil {
		log.Println("Failed Query CreateConversionRate : ", err.Error())
		return err
	}

	return nil
}
