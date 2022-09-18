package route

import (
	conversionHandler "game-currency/src/interface/rest/handlers/conversion"
	currencyHandler "game-currency/src/interface/rest/handlers/currency"
	"game-currency/src/interface/rest/middleware"

	"net/http"

	"github.com/go-chi/chi/v5"
)

func ConversionAppRouter(ch conversionHandler.ConversionHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.CheckAPWebHeader)

	r.Mount("/", ConversionRouter(ch))

	return r
}

func CurrencyAppRouter(ch currencyHandler.CurrencyHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.CheckAPWebHeader)

	r.Mount("/", CurrencyRouter(ch))

	return r
}
