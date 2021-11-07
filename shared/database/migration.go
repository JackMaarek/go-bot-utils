package database

import (
	"github.com/JackMaarek/go-bot-utils/shared/models"
	log "github.com/sirupsen/logrus"
)

// Migrate executes migrations once the db is connected
func Migrate() {
	log.Info("Executing migrations...")
	err := Db.AutoMigrate(&models.Server{}, &models.Birthday{}, &models.User{}, &models.Role{})
	if err != nil {
		log.Warnf("Could not execute migrations")
	}
}
