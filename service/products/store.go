package products

import (
	"database/sql"
	"fmt"

	"github.com/adeshinafalade/ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	products := make([]types.Product, 0)

	for rows.Next() {
		p, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func (s *Store) GetProductsByID(productIds []int) ([]types.Product, error) {

	placeholders := "$1"
	args := make([]interface{}, len(productIds))
	for i, v := range productIds {
		if i != 0 {
			placeholders = fmt.Sprintf("%v,$%d", placeholders, i+1)
		}
		args[i] = v
	}

	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (%s)", placeholders)

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	products := make([]types.Product, 0)

	for rows.Next() {
		p, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func (s *Store) UpdateProduct(p types.Product) error {
	_, err := s.db.Exec("UPDATE products SET name = $1, price = $2::DECIMAL, image = $3, description = $4, quantity = $5 WHERE id = $6", p.Name, p.Price, p.Image, p.Description, p.Quantity, p.ID)
	return err
}

// do a method to create products

func scanRowIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}
