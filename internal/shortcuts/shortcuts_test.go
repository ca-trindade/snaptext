package shortcuts_test

import (
	"database/sql"
	"snaptext_go/internal/shortcuts"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestAddShortcut(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
        CREATE TABLE shortcuts (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            code TEXT NOT NULL UNIQUE,
            text TEXT NOT NULL
        );
    `)
	if err != nil {
		t.Fatal(err)
	}

	err = shortcuts.AddShortcut(db, "//mail", "mail@mail.com")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	var code, text string
	err = db.QueryRow("SELECT code, text FROM shortcuts WHERE code = ?", "//mail").Scan(&code, &text)
	if err != nil {
		t.Errorf("expected no error fetching shortcut, got %v", err)
	}
	if code != "//mail" || text != "mail@mail.com" {
		t.Errorf("expected code: %s, text: %s; got code: %s, text: %s", "//mail", "mail@mail.com", code, text)
	}

}
