package testdata

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-ozzo/ozzo-dbx"
	_ "github.com/lib/pq" // initialize posgresql for test
	"github.com/qiangxue/golang-restful-starter-kit/app"
)

var (
	DB *dbx.DB
)

func init() {
	// the test may be started from the home directory or a subdirectory
	err := app.LoadConfig("./config", "../config")
	if err != nil {
		panic(err)
	}
	DB, err = dbx.MustOpen("postgres", app.Config.DSN)
	if err != nil {
		panic(err)
	}
}

// ResetDB re-create the database schema and re-populate the initial data using the SQL statements in db.sql.
// This method is mainly used in tests.
func ResetDB() *dbx.DB {
	if err := runSQLFile(DB, getSQLFile()); err != nil {
		panic(fmt.Errorf("Error while initializing test database: %s", err))
	}
	return DB
}

func getSQLFile() string {
	if _, err := os.Stat("testdata/db.sql"); err == nil {
		return "testdata/db.sql"
	}
	return "../testdata/db.sql"
}

func runSQLFile(db *dbx.DB, file string) error {
	s, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	lines := strings.Split(string(s), ";")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if _, err := db.NewQuery(line).Execute(); err != nil {
			fmt.Println(line)
			return err
		}
	}
	return nil
}
