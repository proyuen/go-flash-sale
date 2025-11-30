package service

import (
	"context"
	"time"

	"github.com/proyuen/flashSale/Server/internal/db"
	"github.com/proyuen/flashSale/Server/internal/token"
	"github.com/proyuen/flashSale/Server/internal/util"
)

// Service defines the business logic interface
type Service interface {
	// User methods
	CreateUser(ctx context.Context, req CreateUserRequest) (UserResponse, error)
	LoginUser(ctx context.Context, req LoginUserRequest) (LoginUserResponse, error)
	GetUser(ctx context.Context, username string) (UserResponse, error)

	// Product methods
	CreateProduct(ctx context.Context, req CreateProductRequest) (ProductResponse, error)
	ListProducts(ctx context.Context, req ListProductsRequest) ([]ProductResponse, error)
	GetProduct(ctx context.Context, id int64) (ProductResponse, error)
}

type service struct {
	store      db.Querier
	tokenMaker token.Maker
	config     util.Config
}

// NewService creates a new Service instance
func NewService(config util.Config, store db.Querier, tokenMaker token.Maker) Service {
	return &service{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}
}

// DTOs

type UserResponse struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateUserRequest struct {
	Username string
	Password string
	Email    string
}

type LoginUserRequest struct {
	Username string
	Password string
}

type LoginUserResponse struct {
	AccessToken string
	User        UserResponse
}

type ProductResponse struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image_url"`
	Price       string    `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	Price       string `json:"price"`
}

type ListProductsRequest struct {
	PageID   int32 `json:"page_id"`
	PageSize int32 `json:"page_size"`
}
