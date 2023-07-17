package service

import (
	"context"

	"github.com/rodrigoafernandes/desafio-clean-architecture/internal/infra/grpc/pb"
	"github.com/rodrigoafernandes/desafio-clean-architecture/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	FindOrderUseCase   usecase.FindOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, FindOrderUseCase usecase.FindOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		FindOrderUseCase:   FindOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) FindAllOrder(ctx context.Context, in *pb.FindOrderRequest) (*pb.OrderList, error) {
	dto := usecase.FindOrderInputDTO{
		Sort:  in.Sort,
		Page:  int(in.Page),
		Limit: int(in.Limit),
	}
	orders, err := s.FindOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	var output []*pb.CreateOrderResponse
	for _, order := range orders {
		orderResponse := &pb.CreateOrderResponse{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.Price + order.Tax),
		}
		output = append(output, orderResponse)
	}
	return &pb.OrderList{
		Orders: output,
	}, nil
}
