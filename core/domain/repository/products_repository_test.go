package repository

import (
	"testing"

	"github.com/imgustacabral/open-food-go/core/domain/entity"
)

func TestNewProduct(t *testing.T) {
	product := entity.NewProduct()

	if product.ID == "" {
		t.Error("Product ID must not be empty")
	}

	if product.Status != entity.Published {
		t.Error("Product status must be published")
	}

	if product.ImportedT == 0 {
		t.Error("Product imported time must not be zero")
	}
}

func TestProductDelete(t *testing.T) {
	product := entity.NewProduct()
	product.Delete()

	if product.Status != entity.Trash {
		t.Error("Product status must be trash")
	}
}

func TestProductDraft(t *testing.T) {
	product := entity.NewProduct()
	product.Draft()

	if product.Status != entity.Draft {
		t.Error("Product status must be draft")
	}
}
