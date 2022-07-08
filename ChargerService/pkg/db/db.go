package db

import (
	"charger_service/pkg/models"
	"fmt"
	"log"

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
	addresses = []models.Address{
		{Street: "Bulevar cara Lazara 7", City: "Novi Sad", Country: "Serbia", PostalCode: 21000, Latitude: 45.244630, Longitude: 19.845820},
		{Street: "Bulevar Oslobođenja 119", City: "Novi Sad", Country: "Serbia", PostalCode: 21000, Latitude: 45.245199, Longitude: 19.842804},
		{Street: "Bulevar cara Lazara 119", City: "Novi Sad", Country: "Serbia", PostalCode: 21000, Latitude: 45.270740, Longitude: 19.832450},
	}
)

var (
	plugs = []models.Plug{
		{PricePerHour: "300 RSD", Type: "Type 2", ChargingSpeedPerMinute: "22kW"},
		{PricePerHour: "500 RSD", Type: "Type 2", ChargingSpeedPerMinute: "26kW"},
		{PricePerHour: "350 RSD", Type: "Type 2", ChargingSpeedPerMinute: "22kW"},
	}
)

var (
	chargers = []models.Charger{
		{
			Name:        "Procredit Bank",
			Address:     addresses[0],
			WorkTime:    "07-20",
			Description: "WiFi, Park, Shopping, EV Parking",
			Capacity:    2,
			Rating:      3.5,
			Plugs:       []models.Plug{plugs[0]},
		},
		{
			Name:        "Promenada Shopping Mall",
			Address:     addresses[1],
			WorkTime:    "10-22",
			Description: "WiFi, Food, Shopping, Free Parking",
			Capacity:    4,
			Rating:      4.3,
			Plugs:       []models.Plug{plugs[1]},
		},
		{
			Name:        "JKP Cistoca",
			Address:     addresses[2],
			WorkTime:    "07-16",
			Description: "EV Parking",
			Capacity:    3,
			Rating:      2.5,
			Plugs:       []models.Plug{plugs[2]},
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
	conf.DBName = "chargersDB"
	return conf
}

func DropTables() {
	Db.DropTableIfExists("addresses")
	Db.DropTableIfExists("plugs")
	Db.DropTableIfExists("chargers")
	Db.DropTableIfExists("charger_plugs")
}

func AutoMigrateTables() {
	Db.AutoMigrate(&models.Address{})
	Db.AutoMigrate(&models.Plug{})
	Db.AutoMigrate(&models.Charger{})
}

func InitializeData() {
	for _, address := range addresses {
		Db.Create(&address)
	}

	for _, charger := range chargers {
		Db.Create(&charger)
	}
}
