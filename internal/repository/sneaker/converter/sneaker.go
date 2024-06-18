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
func ToSneakerFromRepoList(s []repoModel.Sneaker) []*model.Sneaker {
	list := make([]*model.Sneaker, 0, len(s))
	for _, i := range s {
		list = append(list, &model.Sneaker{
			ID:        i.ID,
			Title:     i.Title,
			Price:     i.Price,
			ImageUrl:  i.ImageUrl,
			Available: i.Available,
			CreatedAt: i.CreatedAt,
			UpdatedAt: i.UpdatedAt,
		})
	}
}
