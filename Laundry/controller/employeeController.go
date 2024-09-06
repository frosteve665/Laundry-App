package controller

import (
	"database/sql"
	"fmt"
	"laundry/model"
	"laundry/usecase"
	"laundry/utils"
)

func GetAllEmployee(db *sql.DB) {
	Employees, err := usecase.GetAllEmployee(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	viewAllEmployee(Employees)
}

func GetNamaEmployee(db *sql.DB) {
	fmt.Print("Masukan Nama Anda: ")
	nama := utils.Scan()
	Employee, err := usecase.GetNamaEmployee(db, nama)
	if err != nil {
		fmt.Println(err)
		return
	}
	viewEmployee(false, *Employee)
}

func InsertEmployee(db *sql.DB) {
	Employee := &model.EmployeeModel{}
	fmt.Print("Masukan Nama Anda: ")
	Employee.Nama = utils.Scan()
	fmt.Print("Masukan Nomor Handphone Anda: ")
	Employee.PhoneNumber = utils.Scan()
	err := usecase.InsertEmployee(db, *Employee)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Insert berhasil")
}

func UpdateEmployee(db *sql.DB) {
	Employee := &model.EmployeeModel{}
	fmt.Print("Masukan Nama Yang Ingin Di Update: ")
	Employee.Nama = utils.Scan()
	fmt.Print("Masukan Nomor Handphone Anda: ")
	Employee.PhoneNumber = utils.Scan()
	err := usecase.UpdateEmployee(db, *Employee)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Update berhasil")
}

func DeleteEmployee(db *sql.DB) {
	fmt.Print("Masukan Nama yang ingin di hapus: ")
	nama := utils.Scan()
	err := usecase.DeleteEmployee(db, nama)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Delete berhasil")
}

func viewAllEmployee(Employees []model.EmployeeModel) {
	fmt.Println("Id\tNama\tPhone Number")
	for _, value := range Employees {
		viewEmployee(true, value)
	}
}

func viewEmployee(index bool, e model.EmployeeModel) {

	if index {
		fmt.Printf("%d\t%s\t%s\n", e.Id, e.Nama, e.PhoneNumber)
	} else {
		fmt.Println("Id\t\t: ", e.Id)
		fmt.Println("Nama\t\t: ", e.Nama)
		fmt.Println("PhoneNumber\t: ", e.PhoneNumber)
	}

}
