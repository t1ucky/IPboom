package util_test

import (
	"tt-crack/util"

	"testing"
)

func TestReadIpList(t *testing.T) {
	ipList := "/tmp/iplist.txt"
	t.Log(util.ReadIpList(ipList))
}

func TestReadUserDict(t *testing.T) {
	userDict := "/tmp/user.dic"
	t.Log(util.ReadUserDict(userDict))
}

func TestReadPasswordDict(t *testing.T) {
	passDict := "/tmp/pass.dic"
	t.Log(util.ReadPasswordDict(passDict))
}
