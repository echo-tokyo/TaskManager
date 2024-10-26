package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/echo-tokyo/TaskManager/settings"
)


// получение соединения с БД
func GetConnection() (*gorm.DB, error) {
	var connection *gorm.DB
	var err error

	connection, err = gorm.Open(mysql.Open(settings.DSN), &gorm.Config{})
	if err != nil {
		return connection, err
	}

	return connection, nil
}
