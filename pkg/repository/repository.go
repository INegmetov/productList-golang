package repository

import (
	productApp "github.com/inegmetov/productList-golang"
	"github.com/jmoiron/sqlx"
)

type ProductList interface {
	Create(list productApp.ProductList) (int, error)
	GetById(listId int) (productApp.ProductList, error)
}
type ProductItem interface {
	Create(listId int ,item productApp.ProductItem) (int, error)
	GetAll(listId int) ([]productApp.ProductItem, error)
	GetAllItems() ([]productApp.ProductItem, error)
	AddItemToList(listId int, item productApp.ProductItem) error

}

type Repository struct {
	ProductList
	ProductItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		ProductList: NewPoductListPostgres(db),
		ProductItem: NewPoductItemPostgres(db),
	}
}
