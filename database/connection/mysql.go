package connection

import (
	"fmt"
	"go-election/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectDB(config config.DBConfig) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database)

	return gorm.Open(mysql.Open(dsn), nil)
}
