package repository

import (
	"fmt"

	productApp "github.com/inegmetov/productList-golang"
	"github.com/jmoiron/sqlx"
)

type PoductListPostgres struct {
	db *sqlx.DB
}

func NewPoductListPostgres(db *sqlx.DB) *PoductListPostgres {
	return &PoductListPostgres{db: db}
}

func (r *PoductListPostgres) Create(list productApp.ProductList) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title) VALUES ($1) RETURNING id", productListTable)
	row := r.db.QueryRow(query, list.Title)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PoductListPostgres) GetById(listId int) (productApp.ProductList, error) {
	var list productApp.ProductList

	query := fmt.Sprintf(`SELECT plt.id, plt.title FROM %s plt WHERE id = $1`, productListTable)
	err := r.db.Get(&list, query, listId)

	return list, err
}
