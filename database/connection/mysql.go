package connection

import (
	"database/sql"
	"fmt"
	"go-election/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() (db *gorm.DB, sqlDB *sql.DB, err error) {
	config := config.Data.DB
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database)

	db, err = gorm.Open(mysql.Open(dsn), nil)
	if err != nil {
		panic(err.Error())
	}

	sqlDB, err = db.DB()
	sqlDB.SetMaxIdleConns(30)
	sqlDB.SetMaxOpenConns(30)
	sqlDB.SetConnMaxIdleTime(time.Minute * 5)

	return
}
