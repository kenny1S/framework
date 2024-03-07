package mysql

import (
	"2108a/high-five/home/day12/framework/config"
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type MysqlConf struct {
	Host     string `json:"Host"`
	Port     string `json:"Port"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	Database string `json:"Database"`
}
type Val struct {
	Mysql MysqlConf `json:"Mysql"`
}

func InitMysql(servername string) error {

	getConfig, err := config.GetConfig(servername, "DEFAULT_GROUP")
	if err != nil {
		return err
	}
	s := new(Val)
	json.Unmarshal([]byte(getConfig), &s)
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", s.Mysql.Username, s.Mysql.Password, s.Mysql.Host, s.Mysql.Port, s.Mysql.Database)
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}
