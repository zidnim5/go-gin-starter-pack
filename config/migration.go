package config

import (
	"Starter/pkg/entities"
)

func MigrationDB(Status string) string {
	// connect to db
	// ConnectDB()
	// TODO: migrating db || dropping db
	if Status == "migrate" {
		migrateDB()
		return "Migrating Successfully"
	} else if Status == "drop" {
		dropTable()
		return "Drop Table Successfully"
	}

	return "Nothing Do"
}

func migrateDB() {
	DB.AutoMigrate(entities.User{}, entities.Todo{})
}

func dropTable() {
	DB.DropTable(entities.User{}, entities.Todo{})
}
