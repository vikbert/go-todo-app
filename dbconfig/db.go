package dbconfig

import (
	"github.com/vikbert/go-todo-app/model"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

const dbPath string = "data/todos.db"

func ConnectSQLite() *gorm.DB {
	return connectDB(sqlite.Open(dbPath))
}

func ConnectPostgreSQL() *gorm.DB {
	return connectDB(postgres.Open("host=localhost user=manfred password=manfred dbname=manfred port=5432 sslmode=disable"))
}

func connectDB(dialector gorm.Dialector) *gorm.DB {
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	ormDbInstance, _ := db.DB()
	ormDbInstance.SetMaxOpenConns(25)
	ormDbInstance.SetMaxIdleConns(25)
	ormDbInstance.SetConnMaxLifetime(5 * time.Minute)

	log.Println("Running migrations")
	errMigration := db.AutoMigrate(
		&model.Todo{},
	)

	if errMigration != nil {
		log.Fatal("DB migration failed!")
	}

	return db
}
