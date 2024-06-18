package sneaker

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/pgxpool"
	"sneakers/internal/model"
	def "sneakers/internal/repository"
	"sneakers/internal/repository/sneaker/converter"
	repoModel "sneakers/internal/repository/sneaker/model"
)

var _ def.SneakerRepository = (*repository)(nil)

type repository struct {
	pool *pgxpool.Pool
}

//ImageUrl: "/sneakers/1.jpg",

func NewRepository(pool *pgxpool.Pool) *repository {
	r := &repository{pool: pool}
	return r
}
func (r *repository) Create(ctx context.Context, id string, info *model.Sneaker) error {
	_, err := r.pool.Exec(ctx, `insert into sneaker (id, title, price, image_url, is_available, created_at, updated_at) 
values (@id, @title, @price, @image_url, @is_available, @created_at, @updated_at)`, pgx.NamedArgs{
		"id":           id,
		"title":        info.Title,
		"price":        info.Price,
		"is_available": true,
		"created_at":   info.CreatedAt,
		"updated_at":   info.UpdatedAt,
	})
	return err
}

func (r *repository) Get(ctx context.Context, id string) (*model.Sneaker, error) {
	rows, err := r.pool.Query(ctx, "select id, title, price, image_url, is_available, created_at, updated_at from sneaker where id=@id", pgx.NamedArgs{"id": id})
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	sneaker, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[repoModel.Sneaker])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return converter.ToSneakerFromRepo(&sneaker), nil
}

func (r *repository) GetList(ctx context.Context, start, count int) ([]*model.Sneaker, int, error) {
	tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, 0, err
	}
	defer tx.Rollback(ctx)
	var counter int

	if err = tx.QueryRow(ctx, "SELECT count(id) FROM sneaker").Scan(&counter); err != nil {
		return nil, 0, err
	}
	rows, err := tx.Query(ctx, "select id, title, price, image_url, is_available, created_at, updated_at from sneaker where id=@id", pgx.NamedArgs{"id": id})
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	sneaker, err := pgx.CollectRows(rows, pgx.RowToStructByName[repoModel.Sneaker])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, 0, nil
		}
		return nil, 0, err
	}
	if err = tx.Commit(ctx); err != nil {
		return nil, 0, err
	}
	return converter.ToSneakerFromRepoList(sneaker), counter, nil
}
