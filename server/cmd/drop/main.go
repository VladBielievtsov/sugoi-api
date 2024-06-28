package main

import (
	"fmt"
	"log"
	"sugoi-api/db"
	"sugoi-api/types"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := db.CreateDatabase(); err != nil {
		log.Fatal(err)
	}

	clearTable(db.DB, &types.Image{})

	fmt.Println("Database cleared successfully.")
}

func clearTable(db *gorm.DB, model interface{}) {
	if err := db.Unscoped().Where("1 = 1").Delete(model).Error; err != nil {
		panic(err)
	}
}
