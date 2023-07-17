package usecase

import "github.com/rodrigoafernandes/desafio-clean-architecture/internal/entity"

type FindOrderUseCase struct {
	orderRepository entity.OrderRepositoryInterface
}

type FindOrderInputDTO struct {
	Sort  string
	Page  int
	Limit int
}

func NewFindOrderUseCase(OrderRepository entity.OrderRepositoryInterface) *FindOrderUseCase {
	return &FindOrderUseCase{
		orderRepository: OrderRepository,
	}
}

func (f *FindOrderUseCase) Execute(input FindOrderInputDTO) ([]OrderOutputDTO, error) {
	orders, err := f.orderRepository.FindAll(input.Page, input.Limit, input.Sort)
	if err != nil {
		return nil, err
	}
	var dtos []OrderOutputDTO
	for _, order := range orders {
		dto := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.Price + order.Tax,
		}
		dtos = append(dtos, dto)
	}
	return dtos, nil
}
