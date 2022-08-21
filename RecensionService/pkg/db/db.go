package db

import (
	"fmt"
	"log"
	"recension_service/pkg/models"

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
	recensions = []models.Recension{
		{
			Username:  "boksi",
			ChargerId: 2,
			Date:      1664116200000,
			Content:   "Charger is poor and speed is low",
			Rate:      2,
		},
		{
			Username:  "velja",
			ChargerId: 1,
			Date:      1664116200000 + 40*60*1000,
			Content:   "This is the coolest charger I have ever seen.",
			Rate:      5,
		},
		{
			Username:  "boksi",
			ChargerId: 1,
			Date:      1664116200000 + 200*60*1000,
			Content:   "Fine charger. Speed can be better",
			Rate:      4,
		},
		{
			Username:  "boksi",
			ChargerId: 3,
			Date:      1664116200000 + 60*60*1000,
			Content:   "You sons of a glitch, repair this charger idiots",
			Rate:      1,
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
	conf.DBName = "recensionsDB"
	return conf
}

func DropTables() {
	Db.DropTableIfExists("recensions")
}

func AutoMigrateTables() {
	Db.AutoMigrate(&models.Recension{})
}

func InitializeData() {
	for _, recension := range recensions {
		Db.Create(&recension)
	}
}
