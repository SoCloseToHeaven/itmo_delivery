package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"itmo_delivery/model"
)

func MustInitializeDB() *gorm.DB {

	// TODO: move to env var or to config
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		"postgres-db",
		"admin",
		"root",
		"itmo_delivery",
	)

	// GORM logger config
	newLogger := logger.New(
		log.Default(),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	// Open Connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Panic(err)
	}

	// AutoMigration
	if err := db.AutoMigrate(&model.User{}, &model.Order{}); err != nil {
		log.Panic(err)
	}

	return db
}
