package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"template/backend/config"
)

// GetPSQL returns a pointer to a gorm.DB instance
// you still need to run the psql init method
func GetPSQL(cfg *config.Config) *gorm.DB {
	var dbURI = fmt.Sprintf("host=%s user=%s dbname=%s port=%d sslmode=disable password=%s TimeZone=%s",
		cfg.Database.Postgres.Host,
		cfg.Database.Postgres.User,
		cfg.Database.Postgres.Database,
		cfg.Database.Postgres.Port,
		cfg.Database.Postgres.Password,
		cfg.Database.Postgres.Timezone,
	)

	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	return db
}
