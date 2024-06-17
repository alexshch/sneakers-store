package sneaker

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) GetSneakerByID(c echo.Context) error {
	id := c.Param("id")
	if sneaker, err := h.sneakerService.Get(context.Background(), id); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	} else {
		return c.JSON(http.StatusOK, newUserResponse(sneaker))
	}
}

func (h *Handler) GetSneakerList(c echo.Context) error {
	if sneakers, err := h.sneakerService.GetList(context.Background()); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	} else {
		return c.JSON(http.StatusOK, newSneakerListResponse(sneakers))
	}
}

func (h *Handler) GetSome(c echo.Context) error {
	return c.JSON(http.StatusOK, struct {
		Items []struct {
			Name  string `json:"name,omitempty"`
			Value int    `json:"value,omitempty"`
		} `json:"_"`
	}{
		Items: []struct {
			Name  string `json:"name,omitempty"`
			Value int    `json:"value,omitempty"`
		}{
			{"IVan", 4},
			{"Vonctor", 5},
			{"VIctor", 6},
		},
	})

}

func (h *Handler) GetError(c echo.Context) error {
	err := errors.New("my dull error")
	return c.JSON(http.StatusBadRequest, err)
}
