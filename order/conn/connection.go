package conn

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectSQL() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("ORDER_DB_USER"),
		os.Getenv("ORDER_DB_PASSWORD"),
		os.Getenv("ORDER_DB_HOST"),
		os.Getenv("ORDER_DB_PORT"),
		os.Getenv("ORDER_DB_NAME"),
	)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
