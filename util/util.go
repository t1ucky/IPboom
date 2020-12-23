package util

import (
	"gopkg.in/cheggaaa/pb.v2"
	"tt-crack/logger"
	"tt-crack/models"
	"tt-crack/vars"

	"fmt"
	"net"
	"sync"
)

var (
	AliveAddr []models.IpAddr
	mutex     sync.Mutex
)

func init() {
	AliveAddr = make([]models.IpAddr, 0)
}

//现在要对存活的地址进行操作了 这里是吧各个端口协议混在一起了，之后还要在判断然后区分，这里是否需要不判断 直接存在不同的数组里，提高效率
func CheckAlive(iparr []string,protocol []string) ([]models.IpAddr) {
	fmt.Println(protocol)
	logger.Log.Infoln("checking ip active")
	vars.ProcessBarActive = pb.StartNew(len(iparr))
	vars.ProcessBarActive.SetTemplate(`{{ rndcolor "Checking progress: " }} {{  percent . "[%.02f%%]" "[?]"| rndcolor}} {{ counters . "[%s/%s]" "[%s/?]" | rndcolor}} {{ bar . "「" "-" (rnd "ᗧ" "◔" "◕" "◷" ) "•" "」" | rndcolor}}  {{rtime . | rndcolor }}`)

	var wg sync.WaitGroup
	wg.Add(len(iparr))

	for _, addr := range iparr {//iparr 是ip的数组，adrr是每个ip值
		go func(addr string) {
			defer wg.Done()
			check(addr,protocol)//传单个ip值和数据库名进去
		}(addr)
	}
	wg.Wait()
	vars.ProcessBarActive.Finish()

	return AliveAddr
}


func check(ipAddr string,protocol []string) {//传进来一条地址，几个端口数组
	/*if vars.UdpProtocols[ipAddr.Protocol] {
		_, err := net.DialTimeout("udp", fmt.Sprintf("%v:%v", ipAddr.Ip, ipAddr.Port), vars.TimeOut)
		if err == nil {
			alive = true
		}
	} else {*/

	for _,prt :=range protocol{
		port := vars.DbPort[prt]
		//fmt.Println("循环port：",port)
		_, err := net.DialTimeout("tcp", fmt.Sprintf("%v:%v", ipAddr, port), vars.TimeOut)
		//这里是看端口是否开放，没有连接plugins
		if err == nil {
			//alive = true//这里有几个alive值 所以在这里调用save吧
			//fmt.Println(ipAddr," port： ",port,"通")
			SaveAddr(models.IpAddr{Ip: ipAddr,Port: port,Protocol: prt})

		}
	}
	vars.ProcessBarActive.Increment()

}

func SaveAddr(ipAddr models.IpAddr) {

	mutex.Lock()
	AliveAddr = append(AliveAddr, ipAddr)
	mutex.Unlock()

}
