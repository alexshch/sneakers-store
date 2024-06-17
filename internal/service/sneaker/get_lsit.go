package sneaker

import (
	"context"
	"sneakers/internal/model"
)

func (s *service) GetList(ctx context.Context) ([]model.Sneaker, error) {
	sneakers, err := s.sneakerRepository.GetList(ctx)
	if err != nil {
		return nil, err
	}
	return sneakers, nil
}
