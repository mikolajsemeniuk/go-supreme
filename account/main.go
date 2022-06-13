package account

import (
	"encoding/json"
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

func (a *Account) RedisReader(document string) error {
	return json.Unmarshal([]byte(document), a)
}

func (a *Account) RedisWrite() string {
	content, err := json.Marshal(a)
	if err != nil {
		return ""
	}
	return string(content)
}
