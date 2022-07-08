package db

import (
	"fmt"
	"log"
	"user_service/pkg/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Configuration struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
}

var (
	users = []models.User{
		{Email: "admin@example.com", Username: "admin", Password: "$2a$12$IVzSpWrAgeWXUak/RIocxOk2k4mGJUAsOh9/.Dzuhl0LmUrdmT64O", Firstname: "Adminko", Lastname: "Adminic", Role: models.Administrator},
	}
)

var Db *gorm.DB
var err error

func Init() {
	conf := CreateConfiguration()

	connection := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s port=%d",
		conf.Host,
		conf.User,
		conf.DBName,
		conf.Password,
		conf.Port,
	)
	dialect := "postgres"

	Db, err = gorm.Open(dialect, connection)
	if err != nil {
		log.Fatal(err)
	}

	// drop tables if exist
	DropTables()

	// create tables
	AutoMigrateTables()

	// populating db
	InitializeData()
}

func CreateConfiguration() Configuration {
	var conf Configuration
	conf.User = "postgres"
	conf.Password = "admin"
	conf.Host = "localhost"
	conf.Port = 5432
	conf.DBName = "usersDB"
	return conf
}

func DropTables() {
	Db.DropTable("users")
}

func AutoMigrateTables() {
	Db.AutoMigrate(&models.User{})
}

func InitializeData() {
	for _, user := range users {
		Db.Create(&user)
	}
}
