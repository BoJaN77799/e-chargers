package db

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"log"
	"reservation_service/pkg/entities"
	"time"

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
	boksi, _ = uuid.FromString("602ec829-89de-49b4-b96f-873ff8c819af")
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
	vehicleOne, _   = uuid.FromString("20d1789f-218b-4b6a-9731-7a599aa7b53c")
	vehicleTwo, _   = uuid.FromString("c1d6eaed-33ef-4b0c-bd2e-3dfc14f88d29")
	vehicleThree, _ = uuid.FromString("5d16ce02-9a6a-4602-aa6a-25e16016112e")
)

var (
	vehicleOneName   = "Skutercic"
	vehicleTwoName   = "Tesla Model 2"
	vehicleThreeName = "Tesla Model 3"
)

var (
	reservations = []entities.Reservation{
		{
			Id:          uuid.NewV4(),
			UserId:      boksi,
			ChargerId:   chargerTwo,
			ChargerName: chargerTwoName,
			VehicleId:   vehicleOne,
			VehicleName: vehicleOneName,
			DateFrom:    time.UnixMilli(1664116200000),
			DateTo:      time.UnixMilli(1664116200000 + 30*60*1000),
		},
		{
			Id:          uuid.NewV4(),
			UserId:      boksi,
			ChargerId:   chargerOne,
			ChargerName: chargerOneName,
			VehicleId:   vehicleThree,
			VehicleName: vehicleThreeName,
			DateFrom:    time.UnixMilli(1664116200000),
			DateTo:      time.UnixMilli(1664116200000 + 40*60*1000),
		},
		{
			Id:          uuid.NewV4(),
			UserId:      boksi,
			ChargerId:   chargerThree,
			ChargerName: chargerThreeName,
			VehicleId:   vehicleTwo,
			VehicleName: vehicleTwoName,
			DateFrom:    time.UnixMilli(1661085000000),
			DateTo:      time.UnixMilli(1661087400000),
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
	conf.DBName = "reservationsDB"
	return conf
}

func DropTables() {
	Db.DropTableIfExists("addresses")
	Db.DropTableIfExists("reservations")
}

func AutoMigrateTables() {
	Db.AutoMigrate(&entities.Reservation{})
}

func InitializeData() {
	for _, reservation := range reservations {
		Db.Create(&reservation)
	}
}
