package shortcuts

import (
	"database/sql"
	"fmt"
)

type Shortcut struct {
	ID   int
	Code string
	Text string
}

func AddShortcut(db *sql.DB, code, text string) error {
	if code == "" || text == "" {
		return fmt.Errorf("code and text cannot be empty")
	}

	_, err := db.Exec("INSERT INTO shortcuts (code, text) VALUES (?, ?)", code, text)
	return err
}

func ListShortcuts(db *sql.DB) ([]Shortcut, error) {
	rows, err := db.Query("SELECT id, code, text FROM shortcuts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shortcuts []Shortcut
	for rows.Next() {
		var s Shortcut
		if err := rows.Scan(&s.ID, &s.Code, &s.Text); err != nil {
			return nil, err
		}
		shortcuts = append(shortcuts, s)
	}
	return shortcuts, nil
}

func DeleteShortcut(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM shortcuts WHERE id = ?", id)
	return err
}
