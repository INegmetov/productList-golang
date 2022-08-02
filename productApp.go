package productApp

type ProductItem struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Kcal        int    `json:"kcal" db:"kcal"`
}

type ProductList struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title" binding:"required"`
}
