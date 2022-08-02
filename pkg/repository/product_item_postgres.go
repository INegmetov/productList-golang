package repository

import (
	"fmt"

	productApp "github.com/inegmetov/productList-golang"
	"github.com/jmoiron/sqlx"
)

type PoductItemPostgres struct {
	db *sqlx.DB
}

func NewPoductItemPostgres(db *sqlx.DB) *PoductItemPostgres {
	return &PoductItemPostgres{db: db}
}

func (r *PoductItemPostgres) Create(listId int, item productApp.ProductItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description, kcal) VALUES ($1, $2, $3) RETURNING id", productItemsTable)
	row := r.db.QueryRow(createItemQuery, item.Name, item.Description, item.Kcal)
	if err := row.Scan(&itemId); err != nil {
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values($1, $2)", listItemsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}
func (r *PoductItemPostgres) GetAll(listId int) ([]productApp.ProductItem, error) {
	var items []productApp.ProductItem
	query := fmt.Sprintf(`SELECT pit.id, pit.title, pit.description, pit.kcal FROM %s pit INNER JOIN %s lit on lit.item_id = pit.id 
						 WHERE lit.list_id=$1`,
		productItemsTable, listItemsTable)
	if err := r.db.Select(&items, query, listId); err != nil {
		return nil, err
	}

	return items, nil
}
func (r *PoductItemPostgres) GetAllItems() ([]productApp.ProductItem, error) {
	var items []productApp.ProductItem
	query := fmt.Sprintf(`SELECT pit.id, pit.title, pit.description, pit.kcal FROM %s pit`, productItemsTable)

	if err := r.db.Select(&items, query); err != nil {
		return nil, err
	}

	return items, nil

}

func (r *PoductItemPostgres) AddItemToList(listId int, item productApp.ProductItem) error {
	// var id int
	// query := fmt.Sprintf("INSERT INTO %s (item_id, list_id) VALUES ($1, $2) RETURNING id", listItemsTable)
	// row := r.db.QueryRow(query, item.Id, listId)
	// if err := row.Scan(&id); err != nil {
	// 	return err
	// }
	return nil
}
