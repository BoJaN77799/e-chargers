package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	uuid "github.com/satori/go.uuid"
	"log"
	"user_service/pkg/entities"
)

type Configuration struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
}

var (
	admin, _ = uuid.FromString("bada76ea-7bfc-4d4a-9e96-32755c9eb3f7")
	boksi, _ = uuid.FromString("602ec829-89de-49b4-b96f-873ff8c819af")
	bane, _  = uuid.FromString("14e0d726-1b6f-4cdc-a520-b0ae7e812b2a")
)

var (
	users = []entities.User{
		{
			Id:          admin,
			Email:       "admin@example.com",
			Role:        entities.Administrator,
			Firstname:   "Adminko",
			Lastname:    "Adminic",
			Password:    "$2a$12$QxDPTEbONfGxoUvxIx5oW.ge3anxohFaFU3Nq1AMMbyYei0jOY.9i",
			Strikes:     0,
			Banned:      false,
			BannedAt:    0,
			BannedUntil: 0,
		},
		{
			Id:          boksi,
			Email:       "boksimus@example.com",
			Role:        entities.RegisteredUser,
			Firstname:   "Boksi",
			Lastname:    "Mus",
			Password:    "$2a$12$QxDPTEbONfGxoUvxIx5oW.ge3anxohFaFU3Nq1AMMbyYei0jOY.9i",
			Strikes:     0,
			Banned:      false,
			BannedAt:    0,
			BannedUntil: 0,
		},
		{
			Id:          bane,
			Email:       "bane_kralj@example.com",
			Role:        entities.RegisteredUser,
			Firstname:   "Bane",
			Lastname:    "Kralj",
			Password:    "$2a$12$QxDPTEbONfGxoUvxIx5oW.ge3anxohFaFU3Nq1AMMbyYei0jOY.9i",
			Strikes:     3,
			Banned:      true,
			BannedAt:    1661000100000,
			BannedUntil: 1663678500000,
		},
	}
)

var (
	vehicle1, _ = uuid.FromString("20d1789f-218b-4b6a-9731-7a599aa7b53c")
	vehicle2, _ = uuid.FromString("c1d6eaed-33ef-4b0c-bd2e-3dfc14f88d29")
	vehicle3, _ = uuid.FromString("5d16ce02-9a6a-4602-aa6a-25e16016112e")
	vehicle4, _ = uuid.FromString("e1bb84c6-570c-4cf7-9df5-bb5bb0d13e78")
)

var (
	vehicles = []entities.Vehicle{
		{Id: vehicle1, Name: "Skutercic", VehicleType: entities.SCOOTER, UserID: boksi},
		{Id: vehicle2, Name: "Tesla Model 2", VehicleType: entities.CAR, UserID: boksi},
		{Id: vehicle3, Name: "Tesla Model 3", VehicleType: entities.CAR, UserID: boksi},
		{Id: vehicle4, Name: "Tesla Model 4", VehicleType: entities.CAR, UserID: bane},
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

	Db, err = gorm.Open("postgres", connection)
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
	conf.Password = "postgres"
	conf.Host = "localhost"
	conf.Port = 5432
	conf.DBName = "usersDB"
	return conf
}

func DropTables() {
	Db.DropTable("users", "vehicles")
}

func AutoMigrateTables() {
	Db.AutoMigrate(&entities.User{}, &entities.Vehicle{})
}

func InitializeData() {
	for _, user := range users {
		Db.Create(&user)
	}

	for _, vehicle := range vehicles {
		Db.Create(&vehicle)
	}
}
