package util_test

import (
	"tt-crack/util"

	"testing"
)

func TestDivideAsset(t *testing.T) {
	ipList := "/tmp/iplist.txt"
	t.Log(util.DivideAsset(util.ReadIpList(ipList)))
}
