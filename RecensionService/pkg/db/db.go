package db

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"log"
	"recension_service/pkg/entities"
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
	boksiUUID, _ = uuid.FromString("602ec829-89de-49b4-b96f-873ff8c819af")
)

var (
	recensions = []entities.Recension{
		{
			Id:        uuid.NewV4(),
			UserId:    boksiUUID,
			ChargerId: uuid.NewV4(),
			Date:      time.Now(),
			Content:   "Charger is awesome, but speed is low",
			Rate:      2,
			Toxic:     0.05,
			Banned:    false,
		},
		{
			Id:        uuid.NewV4(),
			UserId:    boksiUUID,
			ChargerId: uuid.NewV4(),
			Date:      time.Now().Add(10 * time.Minute),
			Content:   "This is the coolest charger I have ever seen.",
			Rate:      5,
			Toxic:     0.08,
			Banned:    false,
		},
		{
			Id:        uuid.NewV4(),
			UserId:    boksiUUID,
			ChargerId: uuid.NewV4(),
			Date:      time.Now().Add(40 * time.Minute),
			Content:   "Fine charger. Speed can be better",
			Rate:      4,
			Toxic:     0.26,
			Banned:    false,
		},
		{
			Id:        uuid.NewV4(),
			UserId:    boksiUUID,
			ChargerId: uuid.NewV4(),
			Date:      time.Now().Add(time.Hour),
			Content:   "Stupid bastards, fix this charger",
			Rate:      1,
			Toxic:     0.98,
			Banned:    false,
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
	conf.DBName = "recensionsDB"
	return conf
}

func DropTables() {
	Db.DropTableIfExists("recensions")
}

func AutoMigrateTables() {
	Db.AutoMigrate(&entities.Recension{})
}

func InitializeData() {
	for _, recension := range recensions {
		Db.Create(&recension)
	}
}
