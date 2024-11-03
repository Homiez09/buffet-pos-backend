package main

import (
	"log"

	"github.com/cs471-buffetpos/buffet-pos-backend/bootstrap"
	"github.com/cs471-buffetpos/buffet-pos-backend/configs"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
)

func main() {
	cfg := configs.NewConfig()
	db := bootstrap.NewDB(cfg)

	if err := db.AutoMigrate(
		&models.User{},
		&models.Table{},
		&models.Category{},
		&models.Menu{},
		&models.Setting{},
		&models.Invoice{},
		&models.Order{},
		&models.OrderItem{},
	); err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Migration completed")
}
