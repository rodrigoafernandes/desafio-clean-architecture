package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rodrigoafernandes/desafio-clean-architecture/internal/entity"
	"github.com/rodrigoafernandes/desafio-clean-architecture/internal/usecase"
	"github.com/rodrigoafernandes/desafio-clean-architecture/pkg/events"
)

type WebOrderHandler struct {
	EventDispatcher   events.EventDispatcherInterface
	OrderRepository   entity.OrderRepositoryInterface
	OrderCreatedEvent events.EventInterface
}

func NewWebOrderHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreatedEvent events.EventInterface,
) *WebOrderHandler {
	return &WebOrderHandler{
		EventDispatcher:   EventDispatcher,
		OrderRepository:   OrderRepository,
		OrderCreatedEvent: OrderCreatedEvent,
	}
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder := usecase.NewCreateOrderUseCase(h.OrderRepository, h.OrderCreatedEvent, h.EventDispatcher)
	output, err := createOrder.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebOrderHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	var input usecase.FindOrderInputDTO
	query := r.URL.Query()
	if query.Get("sort") != "" {
		input.Sort = query.Get("sort")
	}
	if query.Get("page") != "" {
		page, err := strconv.Atoi(query.Get("page"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		input.Page = page
	}
	if query.Get("limit") != "" {
		limit, err := strconv.Atoi(query.Get("limit"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		input.Limit = limit
	}
	w.Header().Set("Content-type", "application/json")
	findOrder := usecase.NewFindOrderUseCase(h.OrderRepository)
	orders, err := findOrder.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(orders) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if err = json.NewEncoder(w).Encode(orders); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
