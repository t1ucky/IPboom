package util

import (
	"fmt"
	"net"
)

var list []string

func incIP(ip net.IP) {

	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func Iplist(cidr string) ([]string,error) {
	ip, ipNet, err := net.ParseCIDR(cidr)

	if err != nil {
		ip = net.ParseIP(cidr)
		list = append(list, ip.String())
		return list,err
	}else {
		fmt.Print(err)
	}
	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); incIP(ip) {
		list = append(list, ip.String())
	}
	return list,err
}
