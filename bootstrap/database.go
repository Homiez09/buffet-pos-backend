package bootstrap

import (
	"fmt"
	"log"

	"github.com/cs471-buffetpos/buffet-pos-backend/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(cfg *configs.Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
