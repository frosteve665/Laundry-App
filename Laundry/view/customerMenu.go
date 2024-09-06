package view

import (
	"database/sql"
	"fmt"
	"laundry/controller"
)

func CustomerMenu(db *sql.DB) {
	for {
		fmt.Println("❆  Menu Customer  ❆")
		fmt.Println("1. View Customer")
		fmt.Println("2. Insert Customer")
		fmt.Println("3. Update Customer")
		fmt.Println("4. Delete Customer")
		fmt.Println("0. Exit")

		var selecterMenu int
		fmt.Print("Select Menu: ")
		fmt.Scanln(&selecterMenu)

		switch selecterMenu {
		case 1:
			fmt.Print("\033[H\033[2J")
			selectViewCustomer(db)
		case 2:
			fmt.Print("\033[H\033[2J")
			controller.InsertCustomer(db)
		case 3:
			fmt.Print("\033[H\033[2J")
			controller.UpdateCustomer(db)
		case 4:
			fmt.Print("\033[H\033[2J")
			controller.DeleteCustomer(db)
		case 0:
			fmt.Print("\033[H\033[2J")
			return
		default:
			fmt.Println("Menu Yang Dipilih Tidak Ada")
			fmt.Println()
		}

	}

}

func selectViewCustomer(db *sql.DB) {
	for {
		fmt.Println("❆     View Customer     ❆")
		fmt.Println("1. View All Customer")
		fmt.Println("2. View Customer By Name")
		fmt.Println("0. Exit")

		var selecterMenu int
		fmt.Print("Select View Mode: ")
		fmt.Scanln(&selecterMenu)
		switch selecterMenu {
		case 1:
			fmt.Print("\033[H\033[2J")
			controller.GetAllCustomer(db)
			return
		case 2:
			fmt.Print("\033[H\033[2J")
			controller.GetNamaCustomer(db)
			return
		case 0:
			fmt.Print("\033[H\033[2J")
			return
		default:
			fmt.Println("Menu Yang Dipilih Tidak Ada")
			fmt.Println()
		}
	}

}
