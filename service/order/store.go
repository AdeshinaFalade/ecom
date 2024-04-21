package order

import (
	"database/sql"

	"github.com/adeshinafalade/ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateOrder(order types.Order) (int, error) {
	// res, err := s.db.Exec("INSERT INTO orders (userId, total, status, address) VALUES ($1,$2::DECIMAL,$3,$4) RETURNING id", order.UserID, order.Total, order.Status, order.Address)
	// if err != nil {
	// 	return 0, err
	// }

	// id, err := res.LastInsertId()
	// if err != nil {
	// 	return 0, err
	// }

	// return int(id), nil

	var id int
	query := "INSERT INTO orders (userId, total, status, address) VALUES ($1, $2::DECIMAL, $3, $4) RETURNING id"
	err := s.db.QueryRow(query, order.UserID, order.Total, order.Status, order.Address).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Store) CreateOrderItem(orderItem types.OrderItem) error {
	_, err := s.db.Exec("INSERT INTO order_items (orderId, productId, quantity, price) VALUES ($1,$2,$3,$4)", orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)
	return err
}
