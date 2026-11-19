package healthz

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	r.Get("/", getHandler)
}

func getHandler(res http.ResponseWriter, _ *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Ok"))
}
