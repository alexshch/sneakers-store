package sneaker

import "sneakers/internal/service"

type Handler struct {
	sneakerService service.SneakerService
}

func NewHandler(s service.SneakerService) *Handler {
	return &Handler{sneakerService: s}
}
