package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdateOn    string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
func GetProducts()Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Black Coffee",
		Description: "Nice Black Coffee",
		Price:       2.35,
		SKU:         "sdsd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Brown Coffee",
		Description: "Nice Brown Coffee",
		Price:       5,
		SKU:         "sdsd343",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
	&Product{
		ID:          3,
		Name:        "White Coffee",
		Description: "Nice White Coffee",
		Price:       5,
		SKU:         "adsd343",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
}
