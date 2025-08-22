package migrations

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunDatabaseMigrations(db *sql.DB, dbName string) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to create migrate driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		dbName,
		driver,
	)
	if err != nil {
		log.Fatalf("❌ Failed to create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("❌ Failed to run migrations: %v", err)
	}

	log.Println("✅ Database migrated successfully")
}
