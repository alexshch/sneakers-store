package repository

import (
	"context"
	"sneakers/internal/model"
)

type SneakerRepository interface {
	Create(ctx context.Context, id int, info *model.Sneaker) error
	Get(ctx context.Context, id int) (*model.Sneaker, error)
	GetList(ctx context.Context, start, count int) ([]*model.Sneaker, error)
}
