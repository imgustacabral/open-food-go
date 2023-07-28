package entity

import (
	"time"

	"github.com/google/uuid"
)

type ProductStatus string

const (
	Draft     ProductStatus = "draft"
	Trash     ProductStatus = "trash"
	Published ProductStatus = "published"
)

type Product struct {
	ID              string        `json:"id"`
	Code            string        `json:"code"`
	Status          ProductStatus `json:"status"`
	ImportedT       int64         `json:"imported_t"`
	URl             string        `json:"url"`
	Creator         string        `json:"creator"`
	CreatedT        int64         `json:"created_t"`
	LastModifiedT   int64         `json:"last_modified_t"`
	ProductName     string        `json:"product_name"`
	Quantity        string        `json:"quantity"`
	Brands          string        `json:"brands"`
	Categories      string        `json:"categories"`
	Labels          string        `json:"labels"`
	Cities          string        `json:"cities"`
	PurchasePlaces  string        `json:"purchase_places"`
	Stores          string        `json:"stores"`
	IngredientsText string        `json:"ingredients_text"`
	Traces          string        `json:"traces"`
	ServingSize     string        `json:"serving_size"`
	ServingQuantity string        `json:"serving_quantity"`
	NutriscoreScore int32         `json:"nutriscrore_score"`
	NutriscoreGrade string        `json:"nutriscore_grade"`
	MainCategory    string        `json:"main_category"`
	ImageUrl        string        `json:"image_url"`
}

func NewProduct() *Product {
	return &Product{
		ID:        uuid.New().String(),
		Status:    Published,
		ImportedT: time.Now().Unix(),
	}
}

func (p *Product) Delete() {
	p.Status = Trash
}

func (p *Product) Draft() {
	p.Status = Draft
}
