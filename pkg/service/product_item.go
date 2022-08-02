package service

import (
	productApp "github.com/inegmetov/productList-golang"
	"github.com/inegmetov/productList-golang/pkg/repository"
)

type ProductItemService struct {
	repo repository.ProductItem
	list repository.ProductList
}

func NewProductItemService(repo repository.ProductItem, list repository.ProductList) *ProductItemService {
	return &ProductItemService{repo: repo, list: list}
}

func (s *ProductItemService) Create(listId int, item productApp.ProductItem) (int, error) {
	return s.repo.Create(listId, item)
}
func (s *ProductItemService) GetAll(listId int) ([]productApp.ProductItem, error) {
	return s.repo.GetAll(listId)
}
func (s *ProductItemService) GetAllItems() ([]productApp.ProductItem, error) {
	return s.repo.GetAllItems()
}

func (s *ProductItemService) AddItemToList(listId int, item productApp.ProductItem) error {
	_, err := s.list.GetById(listId)
	if err != nil {
		return err
	}
	return  s.repo.AddItemToList(listId, item)
}
