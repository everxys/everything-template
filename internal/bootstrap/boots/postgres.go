package boots

import (
	"everything-template/internal/vars"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		vars.Config.Postgres.Addr,
		vars.Config.Postgres.Port,
		vars.Config.Postgres.User,
		vars.Config.Postgres.Password,
		vars.Config.Postgres.Dbname,
	)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open PostgreSQL connection: %v", err)
	}

	vars.DB = db
}
