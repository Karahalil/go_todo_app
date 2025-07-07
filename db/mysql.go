//mysql connection setup will be handled here

package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/karahalil/backend-project/config"
)

var DB *sql.DB

func Connect() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", config.DBUser, config.DBPassword, config.DBName)
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("Error pinging the database:", err)
	}
}

func RunSQLFile(filename string) error {
	sqlbytes, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading SQL file: %w", err)
	}
	sqlStmnt := strings.Split(string(sqlbytes), ";")

	for _, stmt := range sqlStmnt {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue // Skip empty statements
		}
		_, err = DB.Exec(stmt)
		if err != nil {
			return fmt.Errorf("error executing SQL statement: %w", err)
		}
	}
	log.Println("SQL file executed successfully")

	return nil
}
