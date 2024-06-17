package sneaker

import (
	"context"
	"sneakers/internal/model"
	"strconv"
)

func (s *service) Get(ctx context.Context, id string) (*model.Sneaker, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, model.SneakerNotFoundError
	}
	sneaker, err := s.sneakerRepository.Get(ctx, idInt)
	if err != nil {
		return nil, err
	}
	if sneaker == nil {
		return nil, model.SneakerNotFoundError
	}
	return sneaker, nil
}
