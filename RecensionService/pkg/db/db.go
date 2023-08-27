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
	boksi, _ = uuid.FromString("602ec829-89de-49b4-b96f-873ff8c819af")
	bane, _  = uuid.FromString("14e0d726-1b6f-4cdc-a520-b0ae7e812b2a")
)

var (
	chargerOne, _   = uuid.FromString("2b223f7f-87e3-4724-bc3c-5d1bb71e88db")
	chargerTwo, _   = uuid.FromString("927ac833-e699-4764-bbc2-543500ca7965")
	chargerThree, _ = uuid.FromString("37e96e47-e764-4d37-a95c-1b267140502f")
)

var (
	recensions = []entities.Recension{
		{
			Id:        uuid.NewV4(),
			UserId:    boksi,
			ChargerId: chargerOne,
			Date:      time.Now(),
			Content:   "Charger is awesome, but speed is low",
			Rate:      2,
			Toxic:     0.05,
			Banned:    false,
		},
		{
			Id:        uuid.NewV4(),
			UserId:    boksi,
			ChargerId: chargerOne,
			Date:      time.Now().Add(10 * time.Minute),
			Content:   "This is the coolest charger I have ever seen.",
			Rate:      5,
			Toxic:     0.08,
			Banned:    false,
		},
		{
			Id:        uuid.NewV4(),
			UserId:    boksi,
			ChargerId: chargerTwo,
			Date:      time.Now().Add(40 * time.Minute),
			Content:   "Fine charger. Speed can be better",
			Rate:      4,
			Toxic:     0.26,
			Banned:    false,
		},
		{
			Id:        uuid.NewV4(),
			UserId:    bane,
			ChargerId: chargerThree,
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
