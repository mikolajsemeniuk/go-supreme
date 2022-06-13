package router

import (
	"net/http"

	"supreme/account"

	"github.com/gorilla/mux"
)

// Make it handlers and use interface instead of account.Account
type HTTPRouter struct {
	Account account.Account
}

func (r *HTTPRouter) Route(listen string) error {
	router := mux.NewRouter()
	router.HandleFunc("/article", r.Account.HandleList).Methods(http.MethodGet)
	router.HandleFunc("/article/{id}", r.Account.HandleRead).Methods(http.MethodGet)
	http.ListenAndServe(listen, router)
	return nil
}
