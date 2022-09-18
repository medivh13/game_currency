package main

import (
	"context"
	"database/sql"

	usecases "game-currency/src/app/use_cases"

	"game-currency/src/infra/config"
	postgres "game-currency/src/infra/persistence/postgress"

	"game-currency/src/interface/rest"

	ms_log "game-currency/src/infra/log"

	conversionUc "game-currency/src/app/use_cases/conversion"
	currencyUc "game-currency/src/app/use_cases/currency"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

//reupdate by Jody 24 Jan 2022
func main() {
	// init context
	ctx := context.Background()

	// read the server environment variables
	conf := config.Make()

	// check is in production mode
	isProd := false
	if conf.App.Environment == "PRODUCTION" {
		isProd = true
	}

	// logger setup
	m := make(map[string]interface{})
	m["env"] = conf.App.Environment
	m["service"] = conf.App.Name
	logger := ms_log.NewLogInstance(
		ms_log.LogName(conf.Log.Name),
		ms_log.IsProduction(isProd),
		ms_log.LogAdditionalFields(m))

	postgresdb, err := postgres.New(conf.SqlDb, logger)

	// gracefully close connection to persistence storage
	defer func(l *logrus.Logger, sqlDB *sql.DB, dbName string) {
		err := sqlDB.Close()
		if err != nil {
			l.Errorf("error closing sql database %s: %s", dbName, err)
		} else {
			l.Printf("sql database %s successfuly closed.", dbName)
		}
	}(logger, postgresdb.Conn.DB, postgresdb.Conn.DriverName())

	conversionRepository := postgres.NewConversionRepository(postgresdb.Conn)
	httpServer, err := rest.New(
		conf.Http,
		isProd,
		logger,
		usecases.AllUseCases{
			ConversUseCase:  conversionUc.NewConversionUseCase(conversionRepository),
			CurrencyUseCase: currencyUc.NewCurrencyUseCase(conversionRepository),
		},
	)
	if err != nil {
		panic(err)
	}
	httpServer.Start(ctx)

}
