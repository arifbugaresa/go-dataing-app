package databases

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DBConnection *gorm.DB

func InitDatabase() {
	// todo use env file (for test only)
	dsn := "root:password@tcp(localhost:3306)/go_dating_app?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	db.AutoMigrate(&User{})

	DBConnection = db
}
