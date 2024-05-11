package dal

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tticket/pkg/conf"
)

var DB *gorm.DB
var err error

func Init() {
	dsn := "%s:%s@tcp(%s:%s)/tticket?charset=utf8mb4&parseTime=True&loc=Local"
	dsn = fmt.Sprintf(dsn, conf.Config.Mysql.UserName, conf.Config.Mysql.Password, conf.Config.Mysql.Host, conf.Config.Mysql.Port)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
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
