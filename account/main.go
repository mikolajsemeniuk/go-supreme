package account

import (
	"net/http"
	"supreme/order"
	"time"
)

type Account struct {
	Id        string      `json:"id"`
	Email     string      `json:"email"`
	Years     int         `json:"years"`
	Available time.Time   `json:"available"`
	Orders    order.Order `json:"orders"`
}

func (a *Account) HandleList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("accounts"))
}

func (a *Account) HandleRead(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("account"))
}
