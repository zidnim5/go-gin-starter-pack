package databases

import (
	"Starter/config"
	"Starter/src/entities"
)

func MigrationDB(Status string) string {
	// connect to db
	config.ConnectDB()
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
	config.DB.AutoMigrate(entities.User{}, entities.Todo{})
}

func dropTable() {
	config.DB.DropTable(entities.User{}, entities.Todo{})
}
