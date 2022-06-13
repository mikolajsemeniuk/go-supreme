package account

import (
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
