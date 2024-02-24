package order_repo

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ipkalid/order-api/model"
	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	Client *redis.Client
}

func orderIDKey(id uint64) string {
	return fmt.Sprintf("order:%d", id)
}

func (r *RedisRepo) Insert(ctx context.Context, order model.Order) error {
	data, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("failed to encode order: %w", err)
	}

	key := orderIDKey(order.OrderId)

	res := r.Client.SetNX(ctx, key, string(data), 0)

	if err := res.Err(); err != nil {
		return fmt.Errorf("failed to add order: %w", err)
	}
	return nil
}

func (r *RedisRepo) FindByID(ctx context.Context, id uint64) (model.Order, error) {
	key := orderIDKey(id)

	value, err := r.Client.Get(ctx, key).Result()

	if err != nil {
		return model.Order{}, fmt.Errorf("failed to add order: %w", err)
	}
	var order model.Order
	err = json.Unmarshal([]byte(value), &order)
	if err != nil {
		return model.Order{}, fmt.Errorf("failed to decode order: %w", err)
	}
	return order, nil
}

func (r *RedisRepo) DeleteByID(ctx context.Context, id uint64) error {
	key := orderIDKey(id)

	res := r.Client.Del(ctx, key)
	if err := res.Err(); err != nil {
		return fmt.Errorf("failed to delete order: %w", err)
	}
	return nil
}

func (r *RedisRepo) Update(ctx context.Context, order model.Order) error {
	data, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("failed to encode order: %w", err)
	}

	key := orderIDKey(order.OrderId)

	res := r.Client.SetXX(ctx, key, string(data), 0)

	if err := res.Err(); err != nil {
		return fmt.Errorf("failed to update order: %w", err)
	}
	return nil
}

func (r *RedisRepo) FindAll(ctx context.Context, order model.Order) ([]model.Order, error) {
	return nil, nil

}
