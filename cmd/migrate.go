package main

import (
	"log"

	"github.com/cs471-buffetpos/buffet-pos-backend/bootstrap"
	"github.com/cs471-buffetpos/buffet-pos-backend/configs"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"gorm.io/gorm"
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

	// Seed default settings
	if err := initializeDefaultSettings(db); err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Migration completed")
}

func initializeDefaultSettings(db *gorm.DB) error {
	const pricePerPersonKey = "pricePerPerson"
	const defaultPricePerPerson = "250.00"
	var setting models.Setting
	if err := db.Where("key = ?", pricePerPersonKey).First(&setting).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			setting = models.Setting{
				Key:   pricePerPersonKey,
				Value: defaultPricePerPerson,
			}
			if err := db.Create(&setting).Error; err != nil {
				return err
			}
			log.Printf("✅ Inserted default setting for %s: %s\n", pricePerPersonKey, defaultPricePerPerson)
		} else {
			return err
		}
	}
	return nil
}
