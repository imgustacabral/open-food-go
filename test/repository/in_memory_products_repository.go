package repository

import (
	"sync"

	"github.com/imgustacabral/open-food-go/core/domain/entity"
)

type InMemoryProductsRepository struct {
	mu       sync.RWMutex
	products map[string]*entity.Product
}

func NewInMemoryProductsRepository() *InMemoryProductsRepository {
	return &InMemoryProductsRepository{
		products: make(map[string]*entity.Product),
	}
}

func (r *InMemoryProductsRepository) Save(p *entity.Product) (*entity.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.products[p.Code] = p

	return p, nil
}

func (r *InMemoryProductsRepository) FindByCode(code string) (*entity.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	product, ok := r.products[code]
	if !ok {
		return nil, nil
	}

	return product, nil
}

func (r *InMemoryProductsRepository) FindAll(pageNumber int, pageSize int) ([]*entity.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	products := make([]*entity.Product, 0, len(r.products))

	for _, product := range r.products {
		products = append(products, product)
	}

	start := pageNumber * pageSize
	end := start + pageSize

	if start > len(products) {
		return []*entity.Product{}, nil
	}

	if end > len(products) {
		end = len(products)
	}

	return products[start:end], nil
}
