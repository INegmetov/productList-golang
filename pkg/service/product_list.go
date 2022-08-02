package service

import (
	productApp "github.com/inegmetov/productList-golang"
	"github.com/inegmetov/productList-golang/pkg/repository"
)

type ProductListService struct {
	repo repository.ProductList
}

func NewProductListService(repo repository.ProductList) *ProductListService {
	return &ProductListService{repo: repo}
}

func (s *ProductListService) Create(list productApp.ProductList) (int, error) {
	return s.repo.Create(list)
}

func (s *ProductListService) GetById(listId int) (productApp.ProductList, error) {
	return s.repo.GetById(listId)
}
