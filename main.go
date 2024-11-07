package main

import (
	"fmt"
	"snaptext_go/internal/database"
	"snaptext_go/internal/shortcuts"
)

func main() {
	db := database.InitDB("data/shortcuts.db")
	defer db.Close()

	for {
		var choice int
		fmt.Println("1 - Add new Snap")
		fmt.Println("2 - List Snaps")
		fmt.Println("3 - Delete Snap")
		fmt.Println("4 - Quit")
		fmt.Print("Choose option: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var code, text string
			fmt.Print("Write a Snap code")
			fmt.Scan(&code)
			fmt.Print("Write a Snap text block")
			fmt.Scan(&text)
			err := shortcuts.AddShortcut(db, code, text)
			if err != nil {
				fmt.Println("Error adding Snap code:", err)
			} else {
				fmt.Println("Snaptext added!")
			}
		case 2:
			shortcuts, err := shortcuts.ListShortcuts(db)
			if err != nil {
				fmt.Println("Error getting Snaptexts", err)
			} else {
				fmt.Println("Snaptexts:")
				for _, s := range shortcuts {
					fmt.Printf("ID: %d, Código: %s, Texto: %s\n", s.ID, s.Code, s.Text)
				}
			}

		case 3:
			var id int
			fmt.Print("Write Snap ID to delete: ")
			fmt.Scan(&id)
			err := shortcuts.DeleteShortcut(db, id)
			if err != nil {
				fmt.Println("Error deleting Snap:", err)
			} else {
				fmt.Println("Snap deleted!")
			}
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option, try again.")
		}
	}
}
