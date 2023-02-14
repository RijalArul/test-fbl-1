package databases

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"test-fbl-1/server/entities"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

var DB *gorm.DB

func StartDB() {
	userDB := goDotEnvVariable("DB_USER")
	passDB := goDotEnvVariable("DB_PASS")
	dbName := goDotEnvVariable("DB_NAME")
	host := goDotEnvVariable("DB_PORT")
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", userDB, passDB, host, dbName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&entities.Product{}, &entities.Company{}, &entities.User{})
}

func GetDB() *gorm.DB {
	return DB
}
