package services

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
	"toy-project/common/context"
	common "toy-project/common/services"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DBService struct {
	common.Datastore
	Db         *gorm.DB
	CanMigrate bool
	Once       sync.Once
}

const DB_SVC = "db_svc"

func (svc DBService) Id() string {
	return DB_SVC
}

//Configures the connection params based on the flags provided
func (svc *DBService) Configure(ctx *context.Context) error {
	svc.CanMigrate = false //Default

	return svc.Datastore.Configure(ctx)
}

func (svc *DBService) Start() error {
	var err error
	svc.Once.Do(func() {
		err = svc.Connection()
	})

	if err != nil {
		return err
	}

	return svc.Datastore.Start()
}

func (svc *DBService) Connection() error {

	var err error
	dbType := os.Getenv("DB_TYPE")

	if os.Getenv("DB_MIGRATE") == "true" {
		log.Print("[WARN] DB Migration enabled")
		svc.CanMigrate = true
	}

	switch dbType {
	case "postgres":
		log.Println("Using Postgres Database")
		port := os.Getenv("DB_PORT")
		if port == "" {
			port = "5432"
		}

		ssl := os.Getenv("DB_SSL")
		if ssl == "" {
			ssl = "disable"
		}

		svc.Db, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", os.Getenv("DB_URL"), port, os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_DATABASE"), ssl))
		break
	case "sqlite3":
		log.Println("Using SQLite Database")
		svc.Db, err = gorm.Open("sqlite3", os.Getenv("DB_DATABASE"))
		break
	}

	if err != nil {
		return err
	}

	svc.Db.DB().SetMaxIdleConns(10)
	svc.Db.DB().SetMaxOpenConns(100)
	svc.Db.DB().SetConnMaxLifetime(5 * time.Second)

	svc.Db.LogMode(false)
	if os.Getenv("DB_LOGS") == "true" {
		svc.Db.LogMode(true)
	}

	return nil
}

func (svc *DBService) Migrate() error {
	if !svc.CanMigrate {
		log.Println("Skipping table migration - Reason: flag")
		return nil
	}

	log.Println("Migrating tables")
	return svc.Db.AutoMigrate().Error
}
