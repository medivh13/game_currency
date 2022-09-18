package middleware

import (
	"context"
	"errors"
	"net/http"

	commonError "game-currency/src/infra/errors"
	"game-currency/src/interface/rest/response"
)

const (
	XApiKey string = "x-api-key"
)

type gameCurrencyContextKey int

const (
	CtxlinkajaHeader gameCurrencyContextKey = iota + 1
)

type ContexlinkajaHeader struct {
	ApiKey *string
}

func CheckAPWebHeader(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		apiKey := r.Header.Get(XApiKey)

		if apiKey == "" {
			err := errors.New("ApiKey should exist in header")
			cerr := commonError.NewError(commonError.INVALID_HEADER_X_API_KEY, err)
			cerr.SetSystemMessage(err.Error())

			response.NewResponseClient().HttpError(w, cerr)
			return
		}

		val := ContexlinkajaHeader{
			ApiKey: &apiKey,
		}

		ctx = context.WithValue(ctx, CtxlinkajaHeader, val)

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
