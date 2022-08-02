package service

import (
	"context"
	v1 "iyyzh-kratos/api/realword/v1"
)

func (s *RealwordService) GetOrder(ctx context.Context, in *v1.GetOrderRequest) (*v1.GetOrderReply, error) {
	return s.oc.GetOrder(ctx, in)
}

func (s *RealwordService) CreateOrder(ctx context.Context, in *v1.CreateOrderRequest) (*v1.CreateOrderReply, error) {
	return s.oc.CreateOrder(ctx, in)
}

func (s *RealwordService) UpdateOrder(ctx context.Context, in *v1.UpdateOrderRequest) (*v1.UpdateOrderReply, error) {
	return s.oc.UpdateOrder(ctx, in)
}

func (s *RealwordService) DeleteOrder(ctx context.Context, in *v1.DeleteOrderRequest) (*v1.DeleteOrderReply, error) {
	return s.oc.DeleteOrder(ctx, in)
}
