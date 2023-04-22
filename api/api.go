package api

import (
	"ProjectIdeas/monolith/internal/db"
	"context"
)

type APIer interface {
	CreateUser(ctx context.Context, user User) error
	DeleteUserByID(ctx context.Context, id string) error

	UserByID(ctx context.Context, id string) (User, error)
	UserByEmail(ctx context.Context, email string) (User, error)
	WishlistByUserID(ctx context.Context, id string) ([]Product, error)

	AddressByID(ctx context.Context, id string) (Address, error)
	AddressByUserID(ctx context.Context, id string) ([]Address, error)
	AddAddress(ctx context.Context, address Address) error
	DeleteAddressByID(ctx context.Context, id string) error

	PaymentByID(ctx context.Context, id string) (Payment, error)
	PaymentByUserID(ctx context.Context, id string) ([]Payment, error)
	AddPayment(ctx context.Context, payment Payment) error
	DeletePaymentByID(ctx context.Context, id string) error
}

type ApiOptions struct {
	DB db.DBer
}

type Api struct {
	db db.DBer
}

func New(options ApiOptions) (*Api, error) {
	return &Api{options.DB}, nil
}
