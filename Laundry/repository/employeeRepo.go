package repository

import (
	"database/sql"
	"fmt"
	"laundry/model"
)

func GetAllEmployee(db *sql.DB) ([]model.EmployeeModel, error) {
	getAllEmployee := "SELECT * FROM mst_Employee"

	Employees := []model.EmployeeModel{}

	// Execute querry
	rows, err := db.Query(getAllEmployee)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var e model.EmployeeModel
		err := rows.Scan(&e.Id, &e.Nama, &e.PhoneNumber)
		if err != nil {
			return nil, err
		}
		Employees = append(Employees, e)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return Employees, nil
}

func GetNamaEmployee(db *sql.DB, nama string) (*model.EmployeeModel, error) {
	getNamaEmployee := "SELECT * FROM mst_Employee where nama=$1"
	Employee := &model.EmployeeModel{}

	err := db.QueryRow(getNamaEmployee, nama).Scan(&Employee.Id, &Employee.Nama, &Employee.PhoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("nama tidak di temukan")
		} else {
			return nil, err
		}

	}
	return Employee, err
}

func GetIdEmployee(db *sql.DB, id int) (*model.EmployeeModel, error) {
	getIdEmployee := "SELECT * FROM mst_Employee where id=$1"
	Employee := &model.EmployeeModel{}

	err := db.QueryRow(getIdEmployee, id).Scan(&Employee.Id, &Employee.Nama, &Employee.PhoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("id tidak di temukan")
		} else {
			return nil, err
		}

	}
	return Employee, err
}

func InsertEmployee(db *sql.DB, e model.EmployeeModel) error {
	insertEmployee := "INSERT INTO mst_Employee (nama, phone_number) VALUES ($1,$2)"

	_, err := db.Exec(insertEmployee, e.Nama, e.PhoneNumber)

	if err != nil {
		return err
	}

	return nil
}

func UpdateEmployee(db *sql.DB, e model.EmployeeModel) error {
	UpdateEmployee := "UPDATE mst_Employee SET phone_number=$2 where nama=$1"

	_, err := db.Exec(UpdateEmployee, e.Nama, e.PhoneNumber)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEmployee(db *sql.DB, nama string) error {
	deleteEmployee := "DELETE FROM mst_Employee where nama=$1"
	_, err := db.Exec(deleteEmployee, nama)
	if err != nil {
		return err
	}
	return nil
}
