package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	mysql *gorm.DB
)

func InitMysql(username, password, sock string, db string) error {
	var err error
	mysql, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local", username, password, sock, db))
	if err != nil {
		return err
	}
	//mysql.SetLogger(ZapLogger{})
	mysql.SingularTable(true)
	return nil
}

func OpenMysqlDebug() {
	mysql.LogMode(true)
}

func Mysql() *gorm.DB {
	return mysql
}
