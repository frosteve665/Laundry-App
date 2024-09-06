package view

import (
	"database/sql"
	"fmt"
	"laundry/controller"
	"os"
)

func TransaksiMenu(db *sql.DB) {
	for {
		fmt.Println("❆  Menu Transaksi  ❆")
		fmt.Println("1. View Transaksi")
		fmt.Println("2. Insert Transaksi")
		fmt.Println("0. Exit")

		var selecterMenu int
		fmt.Print("Select Menu: ")
		fmt.Scanln(&selecterMenu)

		switch selecterMenu {
		case 1:
			fmt.Print("\033[H\033[2J")
			selectViewTransaksi(db)
		case 2:
			fmt.Println("Menu 1")
			controller.InsertTransaksi(db)
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Menu Yang Dipilih Tidak Ada")
			fmt.Println()
		}

	}
}

func selectViewTransaksi(db *sql.DB) {
	for {
		fmt.Println("❆     View Transaksi     ❆")
		fmt.Println("1. View All Transaksi")
		fmt.Println("2. View Transaksi By Id Transaksi")
		fmt.Println("3. View Transaksi By Id Customer")
		fmt.Println("4. View Transaksi By Id Employee")
		fmt.Println("5. View Transaksi By Id Service")
		fmt.Println("0. Exit")

		var selecterMenu int
		fmt.Print("Select View Mode: ")
		fmt.Scanln(&selecterMenu)
		switch selecterMenu {
		case 1:
			fmt.Print("\033[H\033[2J")
			controller.GetAllTransaksi(db)
			return
		case 2:
			fmt.Print("\033[H\033[2J")
			controller.GetIdTransaksi(db)
			return
		case 3:
			fmt.Print("\033[H\033[2J")
			controller.GetIdCustTransaksi(db)
			return
		case 4:
			fmt.Print("\033[H\033[2J")
			controller.GetIdEmpTransaksi(db)
			return
		case 5:
			fmt.Print("\033[H\033[2J")
			controller.GetIdSrvTransaksi(db)
			return
		case 0:
			fmt.Print("\033[H\033[2J")
			TransaksiMenu(db)
		default:
			fmt.Println("Menu Yang Dipilih Tidak Ada")
			fmt.Println()
		}
	}

}
