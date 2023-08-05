package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/microsoft/go-mssqldb"
	"log"

	"github.com/ryougi-shiky/COMP90018_Backend/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

type MySQLUserRepository struct {
	db *gorm.DB
}

var server = "db-mysql-comp90018-do-user-14450765-0.b.db.ondigitalocean.com"
var port = 25060
var user = "doadmin"
var password = "AVNS_Ae5PE66XCm2T9FJnUZX"
var database = "defaultdb"

func NewUserRepository() UserRepository {
	// Build connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, password, server, port, database)
	// Create connection pool
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		log.Printf("Error opening database connection: %s\n", err.Error())
	}
	fmt.Printf("Database Connected!")

	// Check if table exist, if not, create one
	// if !db.HasTable(&models.User{}){
	// 	db = db.AutoMigrate(&models.User{})
	// 	if db.Error != nil {
	// 		log.Printf("Error migrating database: %s\n", db.Error.Error())
	// 	}
	// } else {
	// 	log.Println("Database migrated successfully")
	// }

	return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *MySQLUserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
