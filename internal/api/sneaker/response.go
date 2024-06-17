package sneaker

import "sneakers/internal/model"

type Sneaker struct {
	ID       int    `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Price    int    `json:"price,omitempty"`
	ImageUrl string `json:"imageUrl,omitempty"`
}

type sneakerResponse struct {
	Sneaker Sneaker `json:"item"`
}

type sneakerListResponse struct {
	Sneakers []Sneaker `json:"items"`
}

func newUserResponse(u *model.Sneaker) *sneakerResponse {
	r := new(sneakerResponse)
	r.Sneaker.ID = u.ID
	r.Sneaker.Title = u.Title
	r.Sneaker.Price = u.Price
	r.Sneaker.ImageUrl = u.ImageUrl
	return r
}

func newSneakerListResponse(l []model.Sneaker) *sneakerListResponse {
	r := new(sneakerListResponse)
	for _, s := range l {
		r.Sneakers = append(r.Sneakers, Sneaker{
			ID:       s.ID,
			Title:    s.Title,
			Price:    s.Price,
			ImageUrl: s.ImageUrl,
		})
	}
	return r
}
