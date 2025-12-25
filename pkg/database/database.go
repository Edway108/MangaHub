package database

//make a database layer
import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)
func InitDB(path string) *sql.DB{
	db,err := sql.Open("sqlite3",path)
	if err != nil {
		log.Fatal("Cannot Open db",err)
	}
	if err := db.Ping();
	err!=nil {
		log.Fatal("Cannot connect to the db",err)
	}
	return db
}
