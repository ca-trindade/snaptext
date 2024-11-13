package shortcuts

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

// Estrutura para os casos de teste
type valueTest struct {
	code, text string
	expectErr  bool
}

var valuesTest = []valueTest{
	{"//mail", "mail@mail.com", false},     // Inserção válida
	{"", "text", true},                     // Erro esperado: código vazio
	{"//code", "", true},                   // Erro esperado: texto vazio
	{"//mail", "other@mail.com", true},     // Erro esperado: duplicado
}

func TestAddShortcut(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Criar tabela em memória
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

	for _, test := range valuesTest {
		err := AddShortcut(db, test.code, test.text)

		if test.expectErr {
			if err == nil {
				t.Errorf("expected error but got nil for code: %q, text: %q", test.code, test.text)
			}
		} else {
			if err != nil {
				t.Errorf("expected no error but got %v for code: %q, text: %q", err, test.code, test.text)
			}
		}
	}
}
