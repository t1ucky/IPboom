package util_test

import (
	"tt-crack/util"

	"testing"
)

func TestGenerateTask(t *testing.T) {
	ipList := "127.0.0.1/24"
	userDic := "/tmp/user.dic"
	passDic := "/tmp/pass.dic"

	users, _ := util.ReadUserDict(userDic)
	passwords, _ := util.ReadPasswordDict(passDic)

	t.Log(util.GenerateTask(util.ReadIpList(ipList), users, passwords))
}

func TestDistributionTask(t *testing.T) {
	ipList := "/tmp/iplist.txt"
	userDic := "/tmp/user.dic"
	passDic := "/tmp/pass.dic"

	users, _ := util.ReadUserDict(userDic)
	passwords, _ := util.ReadPasswordDict(passDic)

	tasks, _ := util.GenerateTask(util.ReadIpList(ipList), users, passwords)
	util.DistributionTask(tasks)
}
