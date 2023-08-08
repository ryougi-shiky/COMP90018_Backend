package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var server = "db-mysql-comp90018-do-user-14450765-0.b.db.ondigitalocean.com"
var port = 25060
var user = "doadmin"
var password = "AVNS_Ae5PE66XCm2T9FJnUZX"
var database = "defaultdb"

type MySQLUserRepository struct {
	db *gorm.DB
}

// Create a new repository object. Connect to the database.
func ConnectToDB() (*gorm.DB, error) {
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, password, server, port, database)
	// Create connection pool
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		log.Printf("Error opening database connection: %s\n", err.Error())
	}
	fmt.Printf("Database Connected!")

	return db, err
}
