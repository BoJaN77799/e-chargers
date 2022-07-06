package db

import (
	"fmt"
	"log"
	"net/url"
	"user_service/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfiguration struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
}

func Init() *gorm.DB {
	// dbURL := "postgres://pg:admin@localhost:5432/usersDB"

	// dbURL := "host=localhost user=pg password=admin dbname=usersDB port=5432 sslmode=disable"

	// db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// db, err := gorm.Open("postgres", "host=localhost port=5432 user=pg dbname=usersDB password=admin")
	// if err != nil {
	// 	panic("failed to connect database")
	// }

	var conf DBConfiguration
	conf.User = "postgres"
	conf.Password = "admin"
	conf.Host = "localhost"
	conf.Port = 5432
	conf.DBName = "usersDB"

	dsn := url.URL{
		User:     url.UserPassword(conf.User, conf.Password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Path:     conf.DBName,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	// drop tables if exist
	db.Migrator().DropTable(&models.User{})

	// create table
	db.AutoMigrate(&models.User{})

	return db
}
