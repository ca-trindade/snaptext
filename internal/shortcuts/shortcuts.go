package shortcuts

import (
	"fmt"
	"database/sql"
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

	var existingID int
	err := db.QueryRow("SELECT id FROM shortcuts WHERE code = ?", code).Scan(&existingID)
	if err == nil {
		return fmt.Errorf("duplicate code: %s", code)
	} else if err != sql.ErrNoRows {
		return err
	}

	_, err = db.Exec("INSERT INTO shortcuts (code, text) VALUES (?, ?)", code, text)
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


