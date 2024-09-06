package controller

import (
	"database/sql"
	"fmt"
	"laundry/model"
	"laundry/usecase"
	"laundry/utils"
)

func GetAllCustomer(db *sql.DB) {
	customers, err := usecase.GetAllCustomer(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	viewAllCustomer(customers)
}

func GetNamaCustomer(db *sql.DB) {
	fmt.Print("Masukan Nama Anda: ")
	nama := utils.Scan()
	customer, err := usecase.GetNamaCustomer(db, nama)
	if err != nil {
		fmt.Println(err)
		return
	}
	viewCustomer(false, *customer)
}

func InsertCustomer(db *sql.DB) {
	customer := &model.CustomerModel{}
	fmt.Print("Masukan Nama Anda: ")
	customer.Nama = utils.Scan()
	fmt.Print("Masukan Nomor Handphone Anda: ")
	customer.PhoneNumber = utils.Scan()
	err := usecase.InsertCustomer(db, *customer)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Insert berhasil")
}

func UpdateCustomer(db *sql.DB) {
	customer := &model.CustomerModel{}
	fmt.Print("Masukan Nama Yang Ingin Di Update: ")
	customer.Nama = utils.Scan()
	fmt.Print("Masukan Nomor Handphone Anda: ")
	customer.PhoneNumber = utils.Scan()
	err := usecase.UpdateCustomer(db, *customer)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Update berhasil")
}

func DeleteCustomer(db *sql.DB) {
	fmt.Print("Masukan Nama yang ingin di hapus: ")
	nama := utils.Scan()
	err := usecase.DeleteCustomer(db, nama)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Delete berhasil")
}

func viewAllCustomer(customers []model.CustomerModel) {
	fmt.Println("Id\tNama\tPhone Number")
	for _, value := range customers {
		viewCustomer(true, value)
	}
}

func viewCustomer(index bool, c model.CustomerModel) {

	if index {
		fmt.Printf("%d\t%s\t%s\n", c.Id, c.Nama, c.PhoneNumber)
	} else {
		fmt.Println("Id\t\t: ", c.Id)
		fmt.Println("Nama\t\t: ", c.Nama)
		fmt.Println("PhoneNumber\t: ", c.PhoneNumber)
	}

}
