package sneaker

import "github.com/labstack/echo/v4"

func (h *Handler) Register(v1 *echo.Group) {
	sneakers := v1.Group("/items")
	sneakers.GET("", h.GetSneakerList)
	sneakers.GET("/:id", h.GetSneakerByID)
	sneakers.GET("/some", h.GetSome)
	v1.GET("/error", h.GetError)
}
