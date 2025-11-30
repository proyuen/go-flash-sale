package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/proyuen/flashSale/Server/internal/db"
)

func (s *service) CreateProduct(ctx context.Context, req CreateProductRequest) (ProductResponse, error) {
	arg := db.CreateProductParams{
		Name:        req.Name,
		Description: req.Description,
		ImageUrl:    req.ImageUrl,
		Price:       req.Price,
	}

	product, err := s.store.CreateProduct(ctx, arg)
	if err != nil {
		return ProductResponse{}, err
	}

	return newProductResponse(product), nil
}

func (s *service) GetProduct(ctx context.Context, id int64) (ProductResponse, error) {
	product, err := s.store.GetProduct(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return ProductResponse{}, fmt.Errorf("product not found: %w", err)
		}
		return ProductResponse{}, err
	}

	return newProductResponse(product), nil
}

func (s *service) ListProducts(ctx context.Context, req ListProductsRequest) ([]ProductResponse, error) {
	arg := db.ListProductsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	products, err := s.store.ListProducts(ctx, arg)
	if err != nil {
		return nil, err
	}

	rsp := make([]ProductResponse, len(products))
	for i, product := range products {
		rsp[i] = newProductResponse(product)
	}

	return rsp, nil
}

func newProductResponse(product db.Product) ProductResponse {
	return ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		ImageUrl:    product.ImageUrl,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
	}
}
