package model

type Product struct {
	ID          string `json:"id"`
	CategoryID  string `json:"category_id"`
	Title       string `json:"title"`
	Alias       string `json:"alias"`
	Content     string `json:"content"`
	Price       string `json:"price"`
	OldPrice    string `json:"old_price"`
	Status      string `json:"status"`
	Keywords    string `json:"keywords"`
	Description string `json:"description"`
	Hit         string `json:"hit"`
}
