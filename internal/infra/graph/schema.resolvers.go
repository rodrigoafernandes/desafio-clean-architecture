package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"

	"github.com/rodrigoafernandes/desafio-clean-architecture/internal/infra/graph/model"
	"github.com/rodrigoafernandes/desafio-clean-architecture/internal/usecase"
)

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input *model.OrderInput) (*model.Order, error) {
	dto := usecase.OrderInputDTO{
		ID:    input.ID,
		Price: float64(input.Price),
		Tax:   float64(input.Tax),
	}
	output, err := r.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &model.Order{
		ID:         output.ID,
		Price:      float64(output.Price),
		Tax:        float64(output.Tax),
		FinalPrice: float64(output.FinalPrice),
	}, nil
}

// ListOrders is the resolver for the ListOrders field.
func (r *queryResolver) ListOrders(ctx context.Context, input *model.ListOrderInput) ([]*model.Order, error) {
	inputDto := usecase.FindOrderInputDTO{
		Sort:  "",
		Page:  0,
		Limit: 0,
	}
	if input != nil {
		sort := "asc"
		var page, limit = 0, 0
		if input.Sort != nil && *input.Sort != "" && *input.Sort != " " {
			sort = *input.Sort
		}
		if input.Page != nil {
			page = *input.Page
		}
		if input.Limit != nil {
			limit = *input.Limit
		}
		inputDto = usecase.FindOrderInputDTO{
			Sort:  sort,
			Page:  page,
			Limit: limit,
		}
	}
	dtos, err := r.FindOrderUseCase.Execute(inputDto)
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for _, dto := range dtos {
		order := &model.Order{
			ID:         dto.ID,
			Price:      dto.Price,
			Tax:        dto.Tax,
			FinalPrice: dto.Price + dto.Tax,
		}
		orders = append(orders, order)
	}
	return orders, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
