package lib

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	// server = os.Getenv("DB_HOST")
	// //port     = os.Getenv("DB_PORT")
	// port     = 1433
	// user     = os.Getenv("DB_USER")
	// password = os.Getenv("DB_PASSWORD")
	// database = os.Getenv("DB_NAME")

	server   = "10.81.5.54"
	port     = 1433
	user     = "Affiliate"
	password = "Affiliate@546510@Us@er"
	database = "PospDB"
)

func InitialMigration() *sql.DB {

	// Connect to database
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	Conn, err := sql.Open("mssql", connString)
	fmt.Printf("type is ==: %s\n", reflect.TypeOf(Conn))
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	fmt.Printf("Connected!\n")
	// defer Conn.Close()

	return Conn
}
