package app

import (
	"2108a/high-five/home/day12/framework/config"
	"2108a/high-five/home/day12/framework/mysql"
)

func Init(servername, ip string, port int, data ...string) error {
	var err error
	err = config.GetClient(ip, port)
	if err != nil {
		return err
	}
	for _, datum := range data {
		switch datum {
		case "mysql":
			err = mysql.InitMysql(servername)
		}

	}
	return err
}
