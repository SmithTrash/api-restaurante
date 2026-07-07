package models

type Pedido struct {
	ID 	int    `json:"id"`
	Cliente string  `json:"cliente"`
	Status string  `json:"status"`
}