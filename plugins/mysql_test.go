

package plugins_test

import (
	"testing"

	"tt-crack/plugins"
	"tt-crack/models"
)

func TestScanMysql(t *testing.T) {
	service := models.Service{Ip: "192.168.31.252", Port: 3306, Protocol: "mysql", Username: "root", Password: "root"}
	t.Log(plugins.ScanMysql(service))
}
