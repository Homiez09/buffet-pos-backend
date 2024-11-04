package usecases

import (
	"github.com/cs471-buffetpos/buffet-pos-backend/configs"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/repositories"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/responses"
)

type OrderUseCase interface {
	GetOrdersByStatus(status string) ([]responses.OrderDetail, error)
	GetOrdersByTableID(tableID string) ([]responses.OrderDetail, error)
	UpdateOrderStatus(orderID string, status string) error
	CreateOrder(order *requests.UserAddOrderRequest) error
}

type OrderService struct {
	orderRepo     repositories.OrderRepository
	orderItemRepo repositories.OrderItemRepository
	config        *configs.Config
}

func NewOrderService(orderRepo repositories.OrderRepository, orderItemRepo repositories.OrderItemRepository, config *configs.Config) OrderUseCase {
	return &OrderService{
		orderRepo:     orderRepo,
		orderItemRepo: orderItemRepo,
		config:        config,
	}
}

func (s *OrderService) GetOrdersByStatus(status string) ([]responses.OrderDetail, error) {
	orders, err := s.orderRepo.GetOrdersByStatus(status)
	if err != nil {
		return nil, err
	}
	var orderDetails []responses.OrderDetail
	for _, order := range orders {
		orderDetails = append(orderDetails, responses.OrderDetail{
			BaseOrder: responses.BaseOrder{
				ID:        order.ID,
				TableID:   order.TableID,
				Status:    order.Status,
				OrderItem: order.OrderItems,
				CreatedAt: order.CreatedAt.String(),
				UpdatedAt: order.UpdatedAt.String(),
			},
		})
	}
	return orderDetails, nil
}

func (s *OrderService) GetOrdersByTableID(tableID string) ([]responses.OrderDetail, error) {
	orders, err := s.orderRepo.GetOrdersByTableID(tableID)
	if err != nil {
		return nil, err
	}
	var orderDetails []responses.OrderDetail
	for _, order := range orders {
		orderDetails = append(orderDetails, responses.OrderDetail{
			BaseOrder: responses.BaseOrder{
				ID:        order.ID,
				TableID:   order.TableID,
				Status:    order.Status,
				OrderItem: order.OrderItems,
				CreatedAt: order.CreatedAt.String(),
				UpdatedAt: order.UpdatedAt.String(),
			},
		})
	}
	return orderDetails, nil
}

func (s *OrderService) UpdateOrderStatus(orderID string, status string) error {
	return s.orderRepo.UpdateOrderStatus(orderID, status)
}

func (s *OrderService) CreateOrder(order *requests.UserAddOrderRequest) error {
	newOrder, err := s.orderRepo.CreateOrder(order)
	if err != nil {
		return err
	}

	for _, item := range order.OrderItems {
		s.orderItemRepo.CreateOrderItem(&item, newOrder.ID.String())
	}

	return nil
}
