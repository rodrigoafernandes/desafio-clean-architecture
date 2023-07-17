package database

import (
	"database/sql"
	"fmt"

	"github.com/rodrigoafernandes/desafio-clean-architecture/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *OrderRepository) FindAll(page, limit int, sort string) ([]entity.Order, error) {
	if sort == "" || (sort != "desc" && sort != "asc") {
		sort = "asc"
	}
	query := fmt.Sprintf("SELECT id, price, tax, final_price FROM orders ORDER BY id %s", sort)
	if page != 0 && limit != 0 {
		offset := (page - 1) * limit
		query = fmt.Sprintf("%s LIMIT %d OFFSET %d ", query, limit, offset)
	}
	rows, err := r.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orders []entity.Order
	for rows.Next() {
		var order entity.Order
		if err = rows.Scan(&order.ID, &order.Price, &order.Tax, &order.FinalPrice); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
