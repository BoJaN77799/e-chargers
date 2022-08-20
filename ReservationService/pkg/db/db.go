package db

import (
	"fmt"
	"log"
	"reservation_service/pkg/models"

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
	reservations = []models.Reservation{
		{
			Username:    "boksi",
			ChargerId:   2,
			ChargerName: "Promenada Shopping Mall",
			VehicleId:   1,
			VehicleName: "Skutercic",
			DateFrom:    1664116200000,
			DateTo:      1664116200000 + 30*60*1000,
		},
		{
			Username:    "velja",
			ChargerId:   1,
			ChargerName: "Procredit Bank",
			VehicleId:   3,
			VehicleName: "Tesla Model 3",
			DateFrom:    1664116200000,
			DateTo:      1664116200000 + 40*60*1000,
		},
		{
			Username:    "boksi",
			ChargerId:   3,
			ChargerName: "JKP Cistoca",
			VehicleId:   2,
			VehicleName: "Tesla Model 2",
			DateFrom:    1664116200000,
			DateTo:      1664116200000 + 60*60*1000,
		},
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
	conf.DBName = "reservationsDB"
	return conf
}

func DropTables() {
	Db.DropTableIfExists("addresses")
	Db.DropTableIfExists("reservations")
}

func AutoMigrateTables() {
	Db.AutoMigrate(&models.Reservation{})
}

func InitializeData() {
	for _, reservation := range reservations {
		Db.Create(&reservation)
	}
}
