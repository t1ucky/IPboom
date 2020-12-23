
package plugins_test

import (
	"tt-crack/models"
	"tt-crack/plugins"

	"testing"
)

func TestScanSsh(t *testing.T) {
	s := models.Service{Ip: "127.0.0.1", Port: 22, Username: "root", Password: "123456", Protocol: "ssh"}
	t.Log(plugins.ScanSsh(s))
}
