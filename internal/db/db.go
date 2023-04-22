package db

import (
	"context"
)

type DBer interface {
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

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
	Addresses []Address
	Payments  []Payment
	Wishlist  []Product
}

type Address struct {
	ID     string
	Street string
	City   string
	State  string
	Zip    string
}

type Payment struct {
	ID             string
	CardNumber     string
	CardType       string
	ExpirationDate string
}

type Product struct {
	ID   string
	Name string
}
