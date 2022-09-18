package route

import (
	"net/http"

	currencyHandler "game-currency/src/interface/rest/handlers/currency"

	"github.com/go-chi/chi/v5"
)

func CurrencyRouter(h currencyHandler.CurrencyHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Get("/", h.GetCurrencyList)
	r.Post("/", h.CreateCurrency)

	return r
}
