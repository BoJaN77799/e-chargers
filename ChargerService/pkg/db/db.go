package db

import (
	"charger_service/pkg/entities"
	"fmt"
	uuid "github.com/satori/go.uuid"
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
	addresses = []entities.Address{
		{Id: uuid.NewV4(), Street: "Bulevar cara Lazara 7", City: "Novi Sad", Country: "Serbia", PostalCode: 21000, Latitude: 45.244630, Longitude: 19.845820},
		{Id: uuid.NewV4(), Street: "Bulevar Oslobođenja 119", City: "Novi Sad", Country: "Serbia", PostalCode: 21000, Latitude: 45.245199, Longitude: 19.842804},
		{Id: uuid.NewV4(), Street: "Sentandrejski put 24a", City: "Novi Sad", Country: "Serbia", PostalCode: 21000, Latitude: 45.270740, Longitude: 19.832450},
	}
)

var (
	chargerOne, _   = uuid.FromString("2b223f7f-87e3-4724-bc3c-5d1bb71e88db")
	chargerTwo, _   = uuid.FromString("927ac833-e699-4764-bbc2-543500ca7965")
	chargerThree, _ = uuid.FromString("37e96e47-e764-4d37-a95c-1b267140502f")
)

var (
	chargerOneName   = "Procredit Bank"
	chargerTwoName   = "Promenada Shopping Mall"
	chargerThreeName = "JKP Cistoca"
)

var (
	chargers = []entities.Charger{
		{
			Id:                     chargerOne,
			Name:                   chargerOneName,
			Address:                addresses[0],
			WorkTimeFrom:           7,
			WorkTimeTo:             20,
			Description:            "WiFi, Park, Shopping, EV Parking",
			Capacity:               2,
			Rating:                 3.5,
			Plugs:                  "Type 1, Type 2, Type 3",
			PricePerHour:           10,
			ChargingSpeedPerMinute: 22,
		},
		{
			Id:                     chargerTwo,
			Name:                   chargerTwoName,
			Address:                addresses[1],
			WorkTimeFrom:           10,
			WorkTimeTo:             22,
			Description:            "WiFi, Food, Shopping, Free Parking",
			Capacity:               4,
			Rating:                 4.3,
			Plugs:                  "Type 1, Type 2",
			PricePerHour:           15,
			ChargingSpeedPerMinute: 26,
		},
		{
			Id:                     chargerThree,
			Name:                   chargerThreeName,
			Address:                addresses[2],
			WorkTimeFrom:           7,
			WorkTimeTo:             16,
			Description:            "EV Parking",
			Capacity:               3,
			Rating:                 2.5,
			Plugs:                  "Type 1, Type 3",
			PricePerHour:           12,
			ChargingSpeedPerMinute: 18,
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
	conf.Password = "postgres"
	conf.Host = "localhost"
	conf.Port = 5432
	conf.DBName = "chargersDB"
	return conf
}

func DropTables() {
	Db.DropTableIfExists("addresses")
	Db.DropTableIfExists("chargers")
}

func AutoMigrateTables() {
	Db.AutoMigrate(&entities.Address{})
	Db.AutoMigrate(&entities.Charger{})
}

func InitializeData() {
	for _, address := range addresses {
		Db.Create(&address)
	}

	for _, charger := range chargers {
		Db.Create(&charger)
	}
}
