

package plugins

import (
	_ "github.com/netxfly/mysql"

	"tt-crack/models"

	"fmt"
	"database/sql"
)

func ScanMysql(service models.Service) (err error, result models.ScanResult) {
	result.Service = service

	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8", service.Username,
		service.Password, service.Ip, service.Port, "mysql")
	db, err := sql.Open("mysql", dataSourceName)

	if err == nil {
		defer db.Close()
		err = db.Ping()
		if err == nil {
			result.Result = true
			fmt.Println("正确密码：",service.Password)
		}

	}
	return err, result
}
