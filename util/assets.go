
package util

import (
	"tt-crack/models"
)

func DivideAsset(ipList []models.IpAddr) (assetsGroup map[string][]models.IpAddr) {
	assetsGroup = make(map[string][]models.IpAddr)
	for _, addr := range ipList {
		protocol := addr.Protocol
		if _, ok := assetsGroup[protocol]; ok {
			assetsGroup[protocol] = append(assetsGroup[protocol], addr)
		} else {
			assetsGroup[protocol] = make([]models.IpAddr, 0)
			assetsGroup[protocol] = append(assetsGroup[protocol], addr)
		}
	}
	return assetsGroup
}
