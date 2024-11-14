package _struct

/*
las estructuras son un tipo de dato que permite agrupar diferentes tipos de datos en un solo tipo de dato
*/

type Product struct {
	ID       string `json:"ID"`
	Name     string
	Price    float64
	Quantity int
	Tag      Type
}

type Type struct {
	ID          string
	Code        string
	Description string
}

func (p *Product) TotalPrice() float64 {
	return p.Price * float64(p.Quantity)
}
