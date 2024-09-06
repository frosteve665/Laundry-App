package controller

import (
	"database/sql"
	"fmt"
	"laundry/model"
	"laundry/usecase"
	"laundry/utils"
	"strconv"
)

func GetAllTransaksi(db *sql.DB) {
	Transaksis, err := usecase.GetAllTransaksi(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	viewAllTransaksi(Transaksis)
}

func GetIdTransaksi(db *sql.DB) {
	fmt.Print("Masukan Id Transaksi: ")
	id, err := strconv.Atoi(utils.Scan())
	if err != nil {
		fmt.Println("Id berupa angka")
	}
	transaksi, err := usecase.GetIdTransaksi(db, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	viewTransaksi(false, *transaksi)
}

func GetIdCustTransaksi(db *sql.DB) {
	fmt.Print("Masukan Id Customer: ")
	id, err := strconv.Atoi(utils.Scan())
	if err != nil {
		fmt.Println("Id Customer berupa angka")
	}
	transaksi, err := usecase.GetIdCustTransaksi(db, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	viewTransaksi(false, *transaksi)
}

func GetIdEmpTransaksi(db *sql.DB) {
	fmt.Print("Masukan Id Employee: ")
	id, err := strconv.Atoi(utils.Scan())
	if err != nil {
		fmt.Println("Id Employee berupa angka")
	}
	transaksi, err := usecase.GetIdEmpTransaksi(db, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	viewTransaksi(false, *transaksi)
}

func GetIdSrvTransaksi(db *sql.DB) {
	fmt.Print("Masukan Id Service: ")
	id, err := strconv.Atoi(utils.Scan())
	if err != nil {
		fmt.Println("Id Service berupa angka")
	}
	transaksi, err := usecase.GetIdSrvTransaksi(db, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	viewTransaksi(false, *transaksi)
}

func InsertTransaksi(db *sql.DB) {
	t := model.TransaksiModel{}
	fmt.Print("Masukan Id Customer:")
	id, err := strconv.Atoi(utils.Scan())
	if err != nil {
		fmt.Println("Id Customer berupa angka")
		return
	}
	t.IdCust = id
	fmt.Print("Masukan Id Employee:")
	id, err = strconv.Atoi(utils.Scan())
	if err != nil {
		fmt.Println("Id Employee berupa angka")
		return
	}
	t.IdEmp = id
	fmt.Print("Masukan Id Service:")
	id, err = strconv.Atoi(utils.Scan())
	if err != nil {
		fmt.Println("Id Service berupa angka")
		return
	}
	t.IdSrv = id
	fmt.Print("Masukan Jumlah Service:")
	jumlah, err := strconv.Atoi(utils.Scan())
	if err != nil {
		fmt.Println("Jumlah Service berupa angka")
		return
	}
	t.Jumlah = jumlah
	err = usecase.InsertTransaksi(db, t)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Insert transaksi berhasil")
}

func viewAllTransaksi(Transaksis []model.TransaksiModel) {
	fmt.Println("Id\tId Customer\tId Employee\tId Service\tJumlah\tTotal")
	for _, value := range Transaksis {
		viewTransaksi(true, value)
	}
}

func viewTransaksi(index bool, t model.TransaksiModel) {

	if index {
		fmt.Printf("%d\t%d\t\t%d\t\t%d\t\t%d\t%d\n", t.Id, t.IdCust, t.IdEmp, t.IdSrv, t.Jumlah, t.Total)
	} else {
		fmt.Println("Id\t\t: ", t.Id)
		fmt.Println("Id Customer\t: ", t.IdCust)
		fmt.Println("Id Employee\t: ", t.IdEmp)
		fmt.Println("Id Service\t: ", t.IdSrv)
		fmt.Println("Jumlah\t\t: ", t.Jumlah)
		fmt.Println("Total\t\t: ", t.Total)
	}

}
