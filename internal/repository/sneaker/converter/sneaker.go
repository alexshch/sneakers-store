package converter

import (
	"sneakers/internal/model"
	repoModel "sneakers/internal/repository/sneaker/model"
)

func ToSneakerFromRepo(s *repoModel.Sneaker) *model.Sneaker {
	return &model.Sneaker{
		ID:       s.ID,
		Title:    s.Title,
		Price:    s.Price,
		ImageUrl: s.ImageUrl,
	}
}
