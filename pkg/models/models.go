package models

type Car struct {
	Id      int     `json:"id"`
	Marca   string  `json:"marca"`
	Nome    string  `json:"nome"`
	Placa   string  `json:"placa"`
	Preco   float64 `json:"preco"`
	Vendido bool    `json:"vendido"`
}
