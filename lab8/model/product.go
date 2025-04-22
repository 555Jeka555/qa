package model

type Product struct {
	ID          int    `json:"id"`
	CategoryID  int    `json:"category_id"`
	Title       string `json:"title"`
	Alias       string `json:"alias"`
	Content     string `json:"content"`
	Price       int    `json:"price"`
	OldPrice    int    `json:"old_price"`
	Status      int    `json:"status"`
	Keywords    string `json:"keywords"`
	Description string `json:"description"`
	Hit         int    `json:"hit"`
}
