package sqlite

import (
	"ProjectIdeas/monolith/internal/db"
	"context"
	"log"
)

func (D *DB) CreateUser(ctx context.Context, user db.User) (rerr error) {
	stmt := `INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?)`

	tx, err := D.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if rerr != nil {
			_ = tx.Rollback()
		}
	}()

	result, err := tx.Exec(stmt, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Println("Error committing transaction: ", err)
		return err
	}

	rows, _ := result.RowsAffected()
	id, _ := result.LastInsertId()
	if rows == 1 {
		log.Printf("User created successfully with ID: %d", id)
	}
	return nil
}

func (D *DB) DeleteUserByID(ctx context.Context, id string) error {
	stmt := `DELETE FROM users WHERE id = ?`

	tx, err := D.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	result, err := tx.Exec(stmt, id)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Println("Error committing transaction: ", err)
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 1 {
		log.Printf("User deleted successfully with ID: %s", id)
	}
	return nil
}

func (D *DB) UserByID(ctx context.Context, id string) (db.User, error) {
	stmt := `SELECT id, first_name, last_name, email FROM users WHERE id = ?`

	var user db.User
	if err := D.db.QueryRow(stmt, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); err != nil {
		return db.User{}, err
	}

	return user, nil
}

func (D *DB) UserByEmail(ctx context.Context, email string) (db.User, error) {
	stmt := `SELECT id, first_name, last_name, email FROM users WHERE email = ?`

	var user db.User
	if err := D.db.QueryRow(stmt, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); err != nil {
		return db.User{}, err
	}

	return user, nil
}

func (D *DB) WishlistByUserID(ctx context.Context, id string) ([]db.Product, error) {
	stmt := `SELECT p.id, p.name 
			FROM products p INNER JOIN users_wishlist uw 
			ON p.id = uw.productid 
			WHERE uw.userid = ? AND uw.inWishlist = 1`

	rows, err := D.db.Query(stmt, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []db.Product
	for rows.Next() {
		var product db.Product
		if err := rows.Scan(&product.ID, &product.Name); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (D *DB) AddressByID(ctx context.Context, id string) (db.Address, error) {
	stmt := `SELECT id, street, city, state, zip 
			FROM address where id = ?`

	var address db.Address
	if err := D.db.QueryRow(stmt, id).Scan(&address.ID, &address.Street, &address.City, &address.State, &address.Zip); err != nil {
		return db.Address{}, err
	}

	return address, nil
}

func (D *DB) AddressByUserID(ctx context.Context, id string) ([]db.Address, error) {
	stmt := `SELECT a.id, a.street, a.city, a.state, a.zip 
			FROM address a inner join users_address ua
			ON a.id = ua.addressid WHERE ua.userid = ?`

	var addresses []db.Address
	rows, err := D.db.Query(stmt, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var addr db.Address
		if err = rows.Scan(&addr.ID, &addr.Street, &addr.City, &addr.State, &addr.Zip); err != nil {
			return nil, err
		}
		addresses = append(addresses, addr)
	}

	return addresses, nil

}

func (D *DB) AddAddress(ctx context.Context, address db.Address) (rerr error) {
	stmt := `INSERT INTO address (street, city, state, zip)
			 VALUES (?,?,?,?)`

	tx, err := D.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if rerr != nil {
			_ = tx.Rollback()
		}
	}()

	result, err := tx.Exec(stmt, address.Street, address.City, address.State, address.Zip)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Println("Error committing transaction: ", err)
		return err
	}

	cnt, err := result.RowsAffected()
	if err != nil || cnt != 1 {
		return err
	}

	log.Println("User inserted successfully")
	return nil
}

func (D *DB) DeleteAddressByID(ctx context.Context, id string) (rerr error) {
	stmt := `DELETE FROM address WHERE id = ?`

	tx, err := D.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if rerr != nil {
			_ = tx.Rollback()
		}
	}()

	rows, err := tx.Exec(stmt, id)
	if err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}

	cnt, err := rows.RowsAffected()
	if err != nil || cnt != 1 {
		log.Printf("error deleting the address : %s", err)
		return err
	}

	log.Println("address deleted")
	return nil
}

func (D DB) PaymentByID(ctx context.Context, id string) (db.Payment, error) {
	stmt := `SELECT id, cardnumber, type, expiration 
			FROM payment_info WHERE id = ?`

	var paymentInfo db.Payment
	if err := D.db.QueryRow(stmt, id).Scan(&paymentInfo.ID, &paymentInfo.CardNumber, &paymentInfo.CardType, &paymentInfo.ExpirationDate); err != nil {
		return db.Payment{}, err
	}

	return paymentInfo, nil
}

func (D *DB) PaymentByUserID(ctx context.Context, id string) (paymentInfos []db.Payment, rerr error) {
	stmt := `SELECT p.id, p.cardnumber, p.type, p.expiration
 			FROM payment_info p INNER JOIN user_payment_info upi 
 			ON p.id = upi.payment_info_id 
 			WHERE upi.userid = ? AND upi.isActive=1`

	tx, err := D.db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if rerr != nil {
			_ = tx.Rollback()
		}
	}()

	rows, err := tx.Query(stmt, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var paymentInfo db.Payment

		if err := rows.Scan(&paymentInfo.ID, &paymentInfo.CardNumber, &paymentInfo.CardType, &paymentInfo.ExpirationDate); err != nil {
			return nil, err
		}
		paymentInfos = append(paymentInfos, paymentInfo)
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return paymentInfos, nil
}

func (D *DB) AddPayment(ctx context.Context, payment db.Payment) (rerr error) {
	stmt := `INSERT INTO payment_info (cardnumber, type, expiration)
 			VALUES (?,?,?)`

	tx, err := D.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if rerr != nil {
			_ = tx.Rollback()
		}
	}()

	result, err := tx.Exec(stmt, payment.CardNumber, payment.CardType, payment.ExpirationDate)
	if err != nil {
		return err
	}
	cnt, err := result.RowsAffected()
	if err != nil || cnt != 1 {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}

func (D *DB) DeletePaymentByID(ctx context.Context, id string) (rerr error) {
	stmt := `DELETE FROM payment_info WHERE id = ?`

	tx, err := D.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if rerr != nil {
			_ = tx.Rollback()
		}
	}()

	result, err := tx.Exec(stmt, id)
	if err != nil {
		return err
	}
	cnt, err := result.RowsAffected()
	if err != nil || cnt != 1 {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
