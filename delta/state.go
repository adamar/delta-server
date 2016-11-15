package delta

import (
	"github.com/adamar/delta-server/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"os"
)

var DBi = SetupGormDB()

func SetupGormDB() *gorm.DB {

	db_url := os.Getenv("DATABASE_URL")
	if db_url == "" {
		db_url = "postgres://postgres:securepass@127.0.0.1:5432/postgres?sslmode=disable"
	}

	conn, err := gorm.Open("postgres", db_url)
	if err != nil {
		//return nil, err
		panic(err)
	}

	runMigrations(conn)

	return conn

}

func runMigrations(M *gorm.DB) {

	M.AutoMigrate(&models.Event{})

}
