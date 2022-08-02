package service

import (
	productApp "github.com/inegmetov/productList-golang"
	"github.com/inegmetov/productList-golang/pkg/repository"
)

type ProductList interface {
	Create(list productApp.ProductList) (int, error)
	GetById(listId int) (productApp.ProductList, error)

}
type ProductItem interface {
	Create(listId int, item productApp.ProductItem) (int, error)
	GetAll(listId int) ([]productApp.ProductItem, error)
	GetAllItems() ([]productApp.ProductItem, error)
	AddItemToList(listId int, item productApp.ProductItem) error

}

type Service struct {
	ProductList
	ProductItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		ProductList: NewProductListService(repos.ProductList),
		ProductItem: NewProductItemService(repos.ProductItem, repos.ProductList),
	}
}
