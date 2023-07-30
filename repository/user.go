package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/microsoft/go-mssqldb"
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

var server = "comp90018t15g02.database.windows.net"
var port = 1433
var user = "comp90018"
var password = "Chelly0412"
var database = "mobile_app"

func NewUserRepository() UserRepository {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)
	// Create connection pool
	db, err := gorm.Open("mssql", connString)
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

func (r *MySQLUserRepository) CreateUser(user *models.User) error{
	return r.db.Create(user).Error
}

func (r *MySQLUserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}