package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	FindAll(page, limit int, sort string) ([]Order, error)
	// GetTotal() (int, error)
}
