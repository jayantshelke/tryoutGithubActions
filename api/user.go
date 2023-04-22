package api

import (
	"context"
	"log"
)

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

func (a Api) CreateUser(ctx context.Context, user User) error {
	log.Println(`called create user in the api layer`)
	return nil
}

func (a Api) DeleteUserByID(ctx context.Context, id string) error {
	log.Println(`called deleteuserbyid from api layer`)
	return nil
}

func (a Api) UserByID(ctx context.Context, id string) (User, error) {
	log.Println(`called userbyid from api layer`)
	return User{
		ID:        `testID`,
		FirstName: `Jayant`,
		LastName:  `Shelke`,
		Email:     `test@fun.com`,
	}, nil
}

func (a Api) UserByEmail(ctx context.Context, email string) (User, error) {
	log.Println(`called userbyemail from api layer`)
	return User{}, nil
}

func (a Api) WishlistByUserID(ctx context.Context, id string) ([]Product, error) {
	log.Println(`called wishlistbyuserid from api layer`)
	return []Product{}, nil
}

func (a Api) AddressByID(ctx context.Context, id string) (Address, error) {
	log.Println(`called addressbyid from api layer`)
	return Address{}, nil
}

func (a Api) AddressByUserID(ctx context.Context, id string) ([]Address, error) {
	log.Println(`called addressbyuserid from api layer`)
	return []Address{}, nil
}

func (a Api) AddAddress(ctx context.Context, address Address) error {
	log.Println(`called addaddress from api layer`)
	return nil
}

func (a Api) DeleteAddressByID(ctx context.Context, id string) error {
	log.Println(`called deleteaddressbyid from api layer`)
	return nil
}

func (a Api) PaymentByID(ctx context.Context, id string) (Payment, error) {
	log.Println(`called paymentbyid from api layer`)
	return Payment{}, nil
}

func (a Api) PaymentByUserID(ctx context.Context, id string) ([]Payment, error) {
	log.Println(`called paymentbyuserid from api layer`)
	return []Payment{}, nil
}

func (a Api) AddPayment(ctx context.Context, payment Payment) error {
	log.Println(`called addpayment from api layer`)
	return nil
}

func (a Api) DeletePaymentByID(ctx context.Context, id string) error {
	log.Println(`called deletepaymentbyid from api layer`)
	return nil
}
