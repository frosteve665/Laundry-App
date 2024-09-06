package repository

import (
	"database/sql"
	"fmt"
	"laundry/model"
)

func GetAllCustomer(db *sql.DB) ([]model.CustomerModel, error) {
	getAllCustomer := "SELECT * FROM mst_customer"

	customers := []model.CustomerModel{}

	// Execute querry
	rows, err := db.Query(getAllCustomer)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var c model.CustomerModel
		err := rows.Scan(&c.Id, &c.Nama, &c.PhoneNumber)
		if err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func GetNamaCustomer(db *sql.DB, nama string) (*model.CustomerModel, error) {
	getNamaCustomer := "SELECT * FROM mst_customer where nama=$1"
	customer := &model.CustomerModel{}

	err := db.QueryRow(getNamaCustomer, nama).Scan(&customer.Id, &customer.Nama, &customer.PhoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("nama tidak di temukan")
		} else {
			return nil, err
		}

	}
	return customer, err
}

func GetIdCustomer(db *sql.DB, id int) (*model.CustomerModel, error) {
	getIdCustomer := "SELECT * FROM mst_customer where id=$1"
	customer := &model.CustomerModel{}

	err := db.QueryRow(getIdCustomer, id).Scan(&customer.Id, &customer.Nama, &customer.PhoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("id tidak di temukan")
		} else {
			return nil, err
		}

	}
	return customer, err
}

func InsertCustomer(db *sql.DB, c model.CustomerModel) error {
	insertCustomer := "INSERT INTO mst_customer (nama, phone_number) VALUES ($1,$2)"

	_, err := db.Exec(insertCustomer, c.Nama, c.PhoneNumber)

	if err != nil {
		return err
	}

	return nil
}

func UpdateCustomer(db *sql.DB, c model.CustomerModel) error {
	UpdateCustomer := "UPDATE mst_customer SET phone_number=$2 where nama=$1"

	_, err := db.Exec(UpdateCustomer, c.Nama, c.PhoneNumber)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCustomer(db *sql.DB, nama string) error {
	deleteCustomer := "DELETE FROM mst_customer where nama=$1"
	_, err := db.Exec(deleteCustomer, nama)
	if err != nil {
		return err
	}
	return nil
}
