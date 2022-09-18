package route

import (
	"net/http"

	conversionHandler "game-currency/src/interface/rest/handlers/conversion"

	"github.com/go-chi/chi/v5"
)

func ConversionRouter(h conversionHandler.ConversionHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Get("/{idFrom}/{idTo}/{ammount}", h.GetConversionResult)
	r.Post("/", h.CreateConversionRate)

	return r
}
