package dal

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"tticket/pkg/conf"
)

var DB *gorm.DB
var err error

func Init() {
	dsn := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn = fmt.Sprintf(dsn, conf.Config.Mysql.UserName, conf.Config.Mysql.Password, conf.Config.Mysql.Host,
		conf.Config.Mysql.Port, conf.Config.Mysql.DbName)

	file, err := os.OpenFile("./db.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	newLogger := logger.New(
		log.New(file, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Millisecond * 80, // Slow SQL threshold
			LogLevel:                  logger.Info,           // Log level
			IgnoreRecordNotFoundError: true,                  // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,                 // Don't include params in the SQL log
			Colorful:                  false,                 // Disable color
		},
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
}

func Update() error {
	return nil
}
