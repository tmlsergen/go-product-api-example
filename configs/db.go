package configs

import (
	"app/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDb() (*gorm.DB, error) {
	dsn := EnvString("MYSQL_USER") + ":" + EnvString("MYSQL_PASSWORD") + "@tcp(" + EnvString("MYSQL_HOST") + ":3306)/" + EnvString("MYSQL_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(models.User{}, models.Product{}, models.DeliveryOption{})

	if err != nil {
		return nil, err
	}
	
	return db, nil
}
