package controller

import (
	"database/sql"
	"fmt"
	"laundry/model"
	"laundry/usecase"
	"laundry/utils"
	"strconv"
)

func GetAllService(db *sql.DB) {
	Services, err := usecase.GetAllService(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	viewAllService(Services)
}

func GetNamaService(db *sql.DB) {
	fmt.Print("Masukan Nama Service: ")
	nama := utils.Scan()
	Service, err := usecase.GetNamaService(db, nama)
	if err != nil {
		fmt.Println(err)
		return
	}
	viewService(false, *Service)
}

func InsertService(db *sql.DB) {
	Service := &model.ServiceModel{}
	fmt.Print("Masukan Nama Service: ")
	Service.Nama = utils.Scan()
	fmt.Print("Masukan Satuan Service: ")
	Service.Satuan = utils.Scan()
	fmt.Print("Masukan Price Service: ")
	Number, err := strconv.Atoi(utils.Scan())
	if err != nil {
		fmt.Println("Price harus berupa angka ")
	}
	Service.Price = Number
	err = usecase.InsertService(db, *Service)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Insert berhasil")
}

func UpdateService(db *sql.DB) {
	Service := &model.ServiceModel{}
	fmt.Print("Masukan Nama Service Yang Ingin Di Update: ")
	Service.Nama = utils.Scan()
	fmt.Print("Masukan Price Service: ")
	number := utils.Scan()
	intNumber, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("Price harus berupa angka ")
	}
	Service.Price = intNumber
	err = usecase.UpdateService(db, *Service)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Update berhasil")
}

func DeleteService(db *sql.DB) {
	fmt.Print("Masukan Nama Service yang ingin di hapus: ")
	nama := utils.Scan()
	err := usecase.DeleteService(db, nama)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Delete berhasil")
}

func viewAllService(Services []model.ServiceModel) {
	fmt.Println("Id\tNama\tSatuan\tPrice")
	for _, value := range Services {
		viewService(true, value)
	}
}

func viewService(index bool, s model.ServiceModel) {

	if index {
		fmt.Printf("%d\t%s\t%s\t%d\n", s.Id, s.Nama, s.Satuan, s.Price)
	} else {
		fmt.Println("Id\t\t: ", s.Id)
		fmt.Println("Nama\t\t: ", s.Nama)
		fmt.Println("Satuan\t: ", s.Satuan)
		fmt.Println("Price\t: ", s.Price)
	}

}
