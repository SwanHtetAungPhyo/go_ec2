package db

import (
	"fmt"
	"os"

	"github.com/SwanHtetAungPhyo/auth/internal/logger"
	"github.com/SwanHtetAungPhyo/auth/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)



var DB *gorm.DB

func Init(){
	logger.LogInit()

	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASS")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", DBUser, DBPassword, DBHost, DBPort)
	DB, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		logger.Logger.Error(err.Error())
		return
	}
	DB.AutoMigrate(&models.User{})
}