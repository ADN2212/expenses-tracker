package DB

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

var dataBaseDNS string

func InitDataBase() {
	
	envErr := godotenv.Load("../.env")

	if envErr != nil {
		panic(".env file not found")
	}

	dataBaseDNS = os.Getenv("DATA_BASE_DNS")

	db, err := sql.Open("postgres", dataBaseDNS)

	if err != nil {
		panic(err.Error())
	}

	defer (func() {
		fmt.Println("Data Base Successfully Initialized!")
		db.Close()
	})()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	_, dbCreateError := db.Exec(`
		CREATE TABLE IF NOT EXISTS transactions (
    		id SERIAL PRIMARY KEY,
    		description TEXT NOT NULL,
    		amount REAL NOT NULL
		);`)

	if dbCreateError != nil {
		panic(dbCreateError.Error())
	}
}

func createDBConnection() *sql.DB {

	db, err := sql.Open("postgres", dataBaseDNS)

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	return db

}
