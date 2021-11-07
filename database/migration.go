package database

import (
	models2 "github.com/JackMaarek/go-bot-utils/models"
	log "github.com/sirupsen/logrus"
)

// Migrate executes migrations once the db is connected
func Migrate() {
	log.Info("Executing migrations...")
	err := Db.AutoMigrate(&models2.Server{}, &models2.Birthday{}, &models2.User{}, &models2.Role{})
	if err != nil {
		log.Warnf("Could not execute migrations")
	}
}
