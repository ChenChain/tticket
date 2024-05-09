package dal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "my_mysql_driver",
		DSN:        "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local", // data source name, 详情参考：https://github.com/go-sql-driver/mysql#dsn-data-source-name
	}), &gorm.Config{})
	_ = db
	_ = err

}

func Insert(any interface{}) error {
	return nil
}

func Update() error {
	return nil
}

func Delete() {
	return
}

func Find() {

}

func FindAll() {

}
