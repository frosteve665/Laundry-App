package view

import (
	"database/sql"
	"fmt"
	"laundry/controller"
)

func EmployeeMenu(db *sql.DB) {
	for {
		fmt.Println("❆  Menu Employee  ❆")
		fmt.Println("1. View Employee")
		fmt.Println("2. Insert Employee")
		fmt.Println("3. Update Employee")
		fmt.Println("4. Delete Employee")
		fmt.Println("0. Exit")

		var selecterMenu int
		fmt.Print("Select Menu: ")
		fmt.Scanln(&selecterMenu)

		switch selecterMenu {
		case 1:
			fmt.Print("\033[H\033[2J")
			selectViewEmployee(db)
		case 2:
			fmt.Print("\033[H\033[2J")
			controller.InsertEmployee(db)
		case 3:
			fmt.Print("\033[H\033[2J")
			controller.UpdateEmployee(db)
		case 4:
			fmt.Print("\033[H\033[2J")
			controller.DeleteEmployee(db)
		case 0:
			fmt.Print("\033[H\033[2J")
			return
		default:
			fmt.Println("Menu Yang Dipilih Tidak Ada")
			fmt.Println()
		}

	}

}

func selectViewEmployee(db *sql.DB) {
	for {
		fmt.Println("❆     View Employee     ❆")
		fmt.Println("1. View All Employee")
		fmt.Println("2. View Employee By Name")
		fmt.Println("0. Exit")

		var selecterMenu int
		fmt.Print("Select View Mode: ")
		fmt.Scanln(&selecterMenu)
		switch selecterMenu {
		case 1:
			fmt.Print("\033[H\033[2J")
			controller.GetAllEmployee(db)
			return
		case 2:
			fmt.Print("\033[H\033[2J")
			controller.GetNamaEmployee(db)
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
