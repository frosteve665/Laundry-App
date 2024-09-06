package view

import (
	"database/sql"
	"fmt"
	"os"
)

func MainMenu(db *sql.DB) {
	for {
		fmt.Println("❆ ENIGMA LAUNDRY ❆")
		fmt.Println("1. Menu Customer")
		fmt.Println("2. Menu Employee")
		fmt.Println("3. Menu Pelayanan")
		fmt.Println("4. Menu Transaksi")
		fmt.Println("0. Exit")

		var selecterMenu int
		fmt.Print("Select Menu: ")
		fmt.Scanln(&selecterMenu)

		switch selecterMenu {
		case 1:
			fmt.Print("\033[H\033[2J")
			CustomerMenu(db)
		case 2:
			fmt.Print("\033[H\033[2J")
			EmployeeMenu(db)
		case 3:
			fmt.Print("\033[H\033[2J")
			ServiceMenu(db)
		case 4:
			fmt.Print("\033[H\033[2J")
			TransaksiMenu(db)
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Menu Yang Dipilih Tidak Ada")
			fmt.Println()
		}

	}

}
