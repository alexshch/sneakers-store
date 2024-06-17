package service

import (
	"context"
	"sneakers/internal/model"
)

type SneakerService interface {
	Create(ctx context.Context, info *model.Sneaker) error
	Get(ctx context.Context, id string) (*model.Sneaker, error)
	GetList(ctx context.Context) ([]model.Sneaker, error)
}

type FavoriteService interface {
	GetFavorite(userID string) ([]model.Sneaker, error)
	AddFavorite(id string) error
	RemoveFavorite(id string) error
}
