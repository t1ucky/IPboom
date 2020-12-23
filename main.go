package main

import (
	"flag"
	"fmt"
	"gopkg.in/fatih/color.v1"
	"net"
	"os"
	"strings"
	"time"
	"tt-crack/util"
	"tt-crack/vars"

)

func usage() {
	fmt.Fprintf(os.Stderr, `支持协议SSH/SMB/MSSQL/MYSQL.
`)
	flag.PrintDefaults()
}

func main() {
//输入和可以修改的东西
	/*
	1.网段 形如192。168。1。1/24  ——vars.IpList
	2.爆破的用户名字典 默认 user.dic _UserDict
	3.爆破的密码字典 默认pass.dic  ——PassDict
	4.要启用的线程数 默认100 ——ScanNum
	5.输出文件默认存储的位置 默认 res.txt ——ResultFile
	6.输入数据库名 string 可以改成string类型在函数里面加以处理 dbname  ——Protocol []string
	7.数据库map默认存储的文件名 port.txt  ——DbPortDict
	 */
	/*
	现在需要修改的问题
	处理的速度
	单进程不用list，快速处理
	修改写入文件的方法，能加快速度？
	之前写的分map的方法应该可以提升速度
	 */
	h := flag.Bool("h", false, "帮助")
	thread := flag.Int("t", 10, "扫描线程")//ScanNum
	ips := flag.String("ip", "192.168.0.1/24", "要爆破的ip网段")//IpList
	user := flag.String("u", "user.dic", "用户名字典")//UserDict
	pass := flag.String("m", "pass.dic", "密码字典")//PassDict
	dbport := flag.String("d", "port.txt", "端口文件")//DbPortDict
	dbname := flag.String("p", "SMB,MSSQL,MYSQL,SSH", "要扫描的协议")//dbname
	result := flag.String("r", "res.txt", "结果保存的文件夹")//dbname


	flag.Parse()

	if *h {
		usage()
		return
	}
	startTime := time.Now()
	vars.ScanNum=*thread
	vars.UserDict=*user
	vars.PassDict=*pass
	vars.IpList=*ips //网段名
	vars.DbPortDict=*dbport
	vars.Protocol=strings.Split(strings.ToUpper(*dbname),",")
	fmt.Println("端口：",vars.Protocol)
	//之后调用都用vars
	vars.ResultFile=*result

	userDict, uErr := util.ReadUserDict(vars.UserDict)
	passDict, pErr := util.ReadPasswordDict(vars.PassDict)
	dbportMap,dErr :=util.ReadDbPort(vars.DbPortDict)
	vars.DbPort = dbportMap

	var ipList []string//ipList返回一个string数组-》包含每个ip
	var iErr error
	if strings.Contains(vars.IpList,"/"){
		ipList,iErr=util.Iplist(vars.IpList)
	}else {
		ip:=net.ParseIP(vars.IpList)
		ipList = append(ipList,ip.String())
		iErr=nil
	}

	color.Cyan("Number of ip list : %d", len(ipList))
	color.Cyan("Number of username dict : %d", len(userDict))
	color.Cyan("Number of password dict : %d", len(passDict))
	color.Cyan("dbname : %v", vars.Protocol)


	if uErr == nil && pErr == nil && iErr==nil && dErr==nil{
		aliveIpList := util.CheckAlive(ipList,vars.Protocol)//返回值是[]model.IpAddr
		fmt.Println("aliveiplist",aliveIpList)
		fmt.Println(vars.Protocol)
		tasks,_:= util.GenerateTask(aliveIpList, userDict, passDict)
		color.Cyan("Number of all task : %d", len(tasks))
		util.RunTask(tasks)
	} else {
		fmt.Println("Read File Err!")
	}

	endTime := time.Now()
	color.Red("Run Time is : %s\n", endTime.Sub(startTime))

}
