package sneaker

import (
	"context"
	"sneakers/internal/model"
	def "sneakers/internal/repository"
	"sneakers/internal/repository/sneaker/converter"
	repoModel "sneakers/internal/repository/sneaker/model"
	"sync"
)

var _ def.SneakerRepository = (*repository)(nil)

type repository struct {
	data map[int]*repoModel.Sneaker
	m    sync.RWMutex
}

func (r *repository) initRepository() *repository {
	r.data[1] = &repoModel.Sneaker{
		ID:       1,
		Title:    "Кросовки Джилли 1",
		Price:    12312,
		ImageUrl: "/sneakers/1.jpg",
	}
	r.data[2] = &repoModel.Sneaker{
		ID:       2,
		Title:    "Кросовки Джилли 33",
		Price:    34531,
		ImageUrl: "/sneakers/2.jpg",
	}
	r.data[3] = &repoModel.Sneaker{
		ID:       3,
		Title:    "",
		Price:    12313,
		ImageUrl: "/sneakers/3.jpg",
	}
	r.data[4] = &repoModel.Sneaker{
		ID:       4,
		Title:    "Кросовки Джилли 14111",
		Price:    1234,
		ImageUrl: "/sneakers/4.jpg",
	}
	r.data[5] = &repoModel.Sneaker{
		ID:       5,
		Title:    "Кросовки Joul 1",
		Price:    41423,
		ImageUrl: "/sneakers/5.jpg",
	}
	r.data[6] = &repoModel.Sneaker{
		ID:       6,
		Title:    "Кросовки Джилли papirte",
		Price:    1231,
		ImageUrl: "/sneakers/5.jpg",
	}
	r.data[7] = &repoModel.Sneaker{
		ID:       7,
		Title:    "Кросовки Джилли faAD",
		Price:    12423,
		ImageUrl: "/sneakers/7.jpg",
	}
	r.data[8] = &repoModel.Sneaker{
		ID:       8,
		Title:    "Кросовки Джилли 1WEQWED",
		Price:    124342,
		ImageUrl: "/sneakers/8.jpg",
	}
	return r
}

func NewRepository() *repository {
	r := &repository{
		data: make(map[int]*repoModel.Sneaker),
	}
	r.initRepository()
	return r
}
func (r *repository) Create(ctx context.Context, id int, info *model.Sneaker) error {
	r.m.Lock()
	defer r.m.Unlock()
	return nil
}

func (r *repository) Get(ctx context.Context, id int) (*model.Sneaker, error) {
	r.m.Lock()
	defer r.m.Unlock()
	if sn, ok := r.data[id]; !ok {
		return nil, nil
	} else {
		return converter.ToSneakerFromRepo(sn), nil
	}
}

func (r *repository) GetList(ctx context.Context) ([]model.Sneaker, error) {
	r.m.Lock()
	defer r.m.Unlock()
	ret := make([]model.Sneaker, 0, 10)
	for _, s := range r.data {
		ret = append(ret, *converter.ToSneakerFromRepo(s))
	}
	return ret, nil
}
