package usecase

import (
	"database/sql"
	"fmt"
	"laundry/model"
	"laundry/repository"
	"strconv"
)

func GetAllEmployee(db *sql.DB) ([]model.EmployeeModel, error) {
	return repository.GetAllEmployee(db)
}

func GetNamaEmployee(db *sql.DB, nama string) (*model.EmployeeModel, error) {
	return repository.GetNamaEmployee(db, nama)
}

func InsertEmployee(db *sql.DB, e model.EmployeeModel) error {
	_, err := repository.GetNamaEmployee(db, e.Nama)
	if err == nil {
		return fmt.Errorf("nama sudah ada")
	}
	_, err = strconv.Atoi(e.PhoneNumber)
	if err != nil || len(e.PhoneNumber) > 13 {
		return fmt.Errorf("phone number harus berupa angka dan maksimal 13 karater")
	}
	return repository.InsertEmployee(db, e)
}

func UpdateEmployee(db *sql.DB, e model.EmployeeModel) error {
	_, err := repository.GetNamaEmployee(db, e.Nama)
	if err != nil {
		return err
	}
	_, err = strconv.Atoi(e.PhoneNumber)
	if err != nil || len(e.PhoneNumber) > 13 {
		return fmt.Errorf("phone number harus berupa angka dan maksimal 13 karater")
	}
	return repository.UpdateEmployee(db, e)
}

func DeleteEmployee(db *sql.DB, nama string) error {
	_, err := repository.GetNamaEmployee(db, nama)
	if err != nil {
		return err
	}
	return repository.DeleteEmployee(db, nama)
}
