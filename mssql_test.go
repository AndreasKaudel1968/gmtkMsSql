package mssql

import (
	"os"
	"testing"
)

func TestRunProc_Logx(t *testing.T) {
	// Setze Umgebungsvariablen für die Datenbankverbindung (ggf. anpassen)
	os.Setenv("DATABASE_USER", "UserLogX")
	os.Setenv("DATABASE_PASSWORD", "UserLogX")
	os.Setenv("DATABASE_SERVER", "localhost")
	os.Setenv("DATABASE_PORT", "1433")
	os.Setenv("DATABASE_DBNAME", "DW")
	os.Setenv("DATABASE_APPNAME", "TestApp")

	// Beispielparameter für die Prozedur Logx
	var result int
	params := []ProcParam[any]{
		{
			Name:      "task",
			Direction: DirectionInput,
			Value:     "TestTask",
		},
		{
			Name:      "obj",
			Direction: DirectionInput,
			Value:     "TestObjekt",
		},
		{
			Name:      "lvl",
			Direction: DirectionInput,
			Value:     1,
		},
		{
			Name:      "evtID",
			Direction: DirectionInput,
			Value:     12345,
		},
		{
			Name:      "msg",
			Direction: DirectionInput,
			Value:     "Testlogeintrag",
		},
		{
			Name:      "src",
			Direction: DirectionInput,
			Value:     "UnitTest",
		},
		{
			Name:      "removeAfterHours",
			Direction: DirectionInput,
			Value:     60,
		},
		{
			Name:       "myResult",
			Direction:  DirectionOutput,
			OutPointer: &result,
		},
	}

	rows, err := RunProc("Logx", params)
	if err != nil {
		t.Fatalf("RunProc(Logx) Fehler: %v", err)
	}
	defer rows.Close()

	// Optional: Überprüfe, ob der Output-Parameter gesetzt wurde
	// if result == "" {
	// 	t.Errorf("Output-Parameter 'Result' wurde nicht gesetzt")
	// }
	t.Logf("Output-Parameter 'Result': %d", result)

}
