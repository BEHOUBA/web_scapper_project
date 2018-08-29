package models

type Product struct {
	Title   string `json:"title"`
	Price   string `json:"price"`
	Picture string `json:"picture"`
	Link    string `json:"link"`
	Origin  string `json:"origin"`
}
