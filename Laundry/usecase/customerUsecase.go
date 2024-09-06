package usecase

import (
	"database/sql"
	"fmt"
	"laundry/model"
	"laundry/repository"
	"strconv"
)

func GetAllCustomer(db *sql.DB) ([]model.CustomerModel, error) {
	return repository.GetAllCustomer(db)
}

func GetNamaCustomer(db *sql.DB, nama string) (*model.CustomerModel, error) {
	return repository.GetNamaCustomer(db, nama)
}

func InsertCustomer(db *sql.DB, c model.CustomerModel) error {
	_, err := repository.GetNamaCustomer(db, c.Nama)
	if err == nil {
		return fmt.Errorf("nama sudah ada")
	}
	_, err = strconv.Atoi(c.PhoneNumber)
	if err != nil || len(c.PhoneNumber) > 13 {
		return fmt.Errorf("phone number harus berupa angka dan maksimal 13 karater")
	}
	return repository.InsertCustomer(db, c)
}

func UpdateCustomer(db *sql.DB, c model.CustomerModel) error {
	_, err := repository.GetNamaCustomer(db, c.Nama)
	if err != nil {
		return err
	}
	_, err = strconv.Atoi(c.PhoneNumber)
	if err != nil || len(c.PhoneNumber) > 13 {
		return fmt.Errorf("phone number harus berupa angka dan maksimal 13 karater")
	}
	return repository.UpdateCustomer(db, c)
}

func DeleteCustomer(db *sql.DB, nama string) error {
	_, err := repository.GetNamaCustomer(db, nama)
	if err != nil {
		return err
	}
	return repository.DeleteCustomer(db, nama)
}
