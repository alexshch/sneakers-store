package sneaker

import (
	"sneakers/internal/repository"
	def "sneakers/internal/service"
)

var _ def.SneakerService = (*service)(nil)

type service struct {
	sneakerRepository repository.SneakerRepository
}

func NewService(r repository.SneakerRepository) *service {
	s := &service{sneakerRepository: r}
	return s
}
