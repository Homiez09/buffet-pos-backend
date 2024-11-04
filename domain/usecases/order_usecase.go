package usecases

import (
	"context"
	"fmt"

	"github.com/cs471-buffetpos/buffet-pos-backend/configs"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/repositories"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/responses"
)

type OrderUseCase interface {
	GetOrdersByStatus(ctx context.Context, status string) ([]responses.OrderDetail, error)
	GetOrdersByTableID(ctx context.Context, tableID string) ([]responses.OrderDetail, error)
	UpdateOrderStatus(ctx context.Context, orderID string, status string) error
	CreateOrder(ctx context.Context, order *requests.UserAddOrderRequest, tableID string) error
	GetOrderHistory(ctx context.Context, tableID string) ([]responses.OrderDetail, error)
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

func (s *OrderService) GetOrdersByStatus(ctx context.Context, status string) ([]responses.OrderDetail, error) {
	orders, err := s.orderRepo.GetOrdersByStatus(ctx, status)
	if err != nil {
		return nil, err
	}

	orderDetails := make([]responses.OrderDetail, 0)
	for _, order := range orders {
		orderItems, err := s.orderItemRepo.GetOrderItemsByOrderID(ctx, order.ID.String())
		if err != nil {
			return nil, err
		}
		orderDetails = append(orderDetails, responses.OrderDetail{
			BaseOrder: responses.BaseOrder{
				ID:        order.ID,
				TableID:   order.TableID,
				Status:    order.Status,
				OrderItem: orderItems,
				CreatedAt: order.CreatedAt,
				UpdatedAt: order.UpdatedAt,
			},
		})
	}
	return orderDetails, nil
}

func (s *OrderService) GetOrdersByTableID(ctx context.Context, tableID string) ([]responses.OrderDetail, error) {
	orders, err := s.orderRepo.GetOrdersByTableID(ctx, tableID)
	if err != nil {
		return nil, err
	}
	orderDetails := make([]responses.OrderDetail, 0)
	for _, order := range orders {
		orderItems, err := s.orderItemRepo.GetOrderItemsByOrderID(ctx, order.ID.String())
		if err != nil {
			return nil, err
		}
		orderDetails = append(orderDetails, responses.OrderDetail{
			BaseOrder: responses.BaseOrder{
				ID:        order.ID,
				TableID:   order.TableID,
				Status:    order.Status,
				OrderItem: orderItems,
				CreatedAt: order.CreatedAt,
				UpdatedAt: order.UpdatedAt,
			},
		})
	}
	return orderDetails, nil
}

func (s *OrderService) UpdateOrderStatus(ctx context.Context, orderID string, status string) error {
	return s.orderRepo.UpdateOrderStatus(ctx, orderID, status)
}

func (s *OrderService) CreateOrder(ctx context.Context, order *requests.UserAddOrderRequest, tableID string) error {
	newOrder, err := s.orderRepo.CreateOrder(ctx, order, tableID)
	if err != nil {
		return err
	}

	for _, item := range order.OrderItems {
		s.orderItemRepo.CreateOrderItem(ctx, &item, newOrder.ID.String())
	}

	return nil
}

func (s *OrderService) GetOrderHistory(ctx context.Context, tableID string) ([]responses.OrderDetail, error) {
	orders, err := s.orderRepo.GetOrderHistory(ctx, tableID)
	if err != nil {
		return nil, err
	}

	orderDetails := make([]responses.OrderDetail, 0)
	for _, order := range orders {
		orderItems := make([]models.OrderItem, 0)
		orderItem, err := s.orderItemRepo.GetOrderItemsByOrderID(ctx, order.ID.String())
		fmt.Println(orderItem)
		if err != nil {
			return nil, err
		}
		for _, item := range orderItem {
			orderItems = append(orderItems, item)
		}
		orderDetails = append(orderDetails, responses.OrderDetail{
			BaseOrder: responses.BaseOrder{
				ID:        order.ID,
				TableID:   order.TableID,
				Status:    order.Status,
				OrderItem: orderItems,
				CreatedAt: order.CreatedAt,
				UpdatedAt: order.UpdatedAt,
			},
		})
	}
	return orderDetails, nil
}
