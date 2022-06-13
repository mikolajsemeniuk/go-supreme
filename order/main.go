package order

type Order struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Failures int    `json:"failures"`
}
