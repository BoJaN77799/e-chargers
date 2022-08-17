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
		{Email: "admin@example.com", Username: "admin", Password: "$2a$12$QxDPTEbONfGxoUvxIx5oW.ge3anxohFaFU3Nq1AMMbyYei0jOY.9i", Firstname: "Adminko", Lastname: "Adminic", Role: models.Administrator},
		{Email: "boksimus@example.com", Username: "boksi", Password: "$2a$12$QxDPTEbONfGxoUvxIx5oW.ge3anxohFaFU3Nq1AMMbyYei0jOY.9i", Firstname: "Boksi", Lastname: "Mus", Role: models.RegisteredUser},
		{Email: "velja_zajecar@example.com", Username: "velja", Password: "$2a$12$QxDPTEbONfGxoUvxIx5oW.ge3anxohFaFU3Nq1AMMbyYei0jOY.9i", Firstname: "Velja", Lastname: "Tomic", Role: models.RegisteredUser},
	}
)

var (
	vehicles = []models.Vehicle{
		{Name: "Skutercic", VehicleType: models.SCOOTER, UserID: 2},
		{Name: "Tesla Model 2", VehicleType: models.CAR, UserID: 2},
		{Name: "Tesla Model 3", VehicleType: models.CAR, UserID: 3},
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
	Db.DropTable("users", "vehicles")
}

func AutoMigrateTables() {
	Db.AutoMigrate(&models.User{}, &models.Vehicle{})
}

func InitializeData() {
	for _, user := range users {
		Db.Create(&user)
	}

	for _, vehicle := range vehicles {
		Db.Create(&vehicle)
	}
}
