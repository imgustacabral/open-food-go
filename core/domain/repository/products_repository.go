package repository

import "github.com/imgustacabral/open-food-go/core/domain/entity"

type ProductsRepository interface {
	Save(prouct *entity.Product) (*entity.Product, error)
	FindByCode(code string) (*entity.Product, error)
	FindAll(pageNumber int, pageSize int) ([]*entity.Product, error)
}
