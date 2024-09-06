package view

import (
	"database/sql"
	"fmt"
	"laundry/controller"
)

func ServiceMenu(db *sql.DB) {
	for {
		fmt.Println("❆  Menu Service  ❆")
		fmt.Println("1. View Service")
		fmt.Println("2. Insert Service")
		fmt.Println("3. Update Service")
		fmt.Println("4. Delete Service")
		fmt.Println("0. Exit")

		var selecterMenu int
		fmt.Print("Select Menu: ")
		fmt.Scanln(&selecterMenu)

		switch selecterMenu {
		case 1:
			fmt.Print("\033[H\033[2J")
			selectViewService(db)
		case 2:
			fmt.Print("\033[H\033[2J")
			controller.InsertService(db)
		case 3:
			fmt.Print("\033[H\033[2J")
			controller.UpdateService(db)
		case 4:
			fmt.Print("\033[H\033[2J")
			controller.DeleteService(db)
		case 0:
			fmt.Print("\033[H\033[2J")
			return
		default:
			fmt.Println("Menu Yang Dipilih Tidak Ada")
			fmt.Println()
		}

	}

}

func selectViewService(db *sql.DB) {
	for {
		fmt.Println("❆     View Service     ❆")
		fmt.Println("1. View All Service")
		fmt.Println("2. View Service By Name")
		fmt.Println("0. Exit")

		var selecterMenu int
		fmt.Print("Select View Mode: ")
		fmt.Scanln(&selecterMenu)
		switch selecterMenu {
		case 1:
			fmt.Print("\033[H\033[2J")
			controller.GetAllService(db)
			return
		case 2:
			fmt.Print("\033[H\033[2J")
			controller.GetNamaService(db)
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
