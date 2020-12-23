package models

import (
	"github.com/patrickmn/go-cache"
	"tt-crack/logger"
	"tt-crack/util/hash"
	"tt-crack/vars"

	"encoding/gob"
	"fmt"
	"os"
	"time"
)

func init() {
	gob.Register(Service{})
	gob.Register(ScanResult{})
}

func SaveResult(err error, result ScanResult) {
	if err == nil && result.Result {
		var k string
		protocol := result.Service.Protocol

		if protocol == "REDIS" || protocol == "FTP" {
			k = fmt.Sprintf("%v-%v-%v", result.Service.Ip, result.Service.Port, result.Service.Protocol)
		} else {
			k = fmt.Sprintf("%v-%v-%v", result.Service.Ip, result.Service.Port, result.Service.Username)
		}

		h := hash.MakeTaskHash(k)
		hash.SetTaskHask(h)

		_, found := vars.CacheService.Get(k)
		if !found {
			logger.Log.Infof("Ip: %v, Port: %v, Protocol: [%v], Username: %v, Password: %v", result.Service.Ip,
				result.Service.Port, result.Service.Protocol, result.Service.Username, result.Service.Password)
				fmt.Println(result.Service)
		}
		vars.CacheService.Set(k, result, cache.NoExpiration)
		Dump()
	}
}

func SavaResultToFile() (error) {
	return vars.CacheService.SaveFile("res.txt")
}

func CacheStatus() (count int, items map[string]cache.Item) {
	count = vars.CacheService.ItemCount()
	items = vars.CacheService.Items()
	return count, items
}

func ResultTotal() {
	vars.ProgressBar.Finish()
	logger.Log.Info(fmt.Sprintf("Finshed scan, total result: %v, used time: %v",
		vars.CacheService.ItemCount(),
		time.Since(vars.StartTime)))
}

func LoadResultFromFile() {
	vars.CacheService.LoadFile(vars.ResultFile)
	vars.ProgressBar.Finish()
	logger.Log.Info(fmt.Sprintf("Finshed scan, total result: %v", vars.CacheService.ItemCount()))
}

func Dump() (err error) {
	LoadResultFromFile()
	err = DumpToFile(vars.ResultFile)
	if err != nil {
		logger.Log.Fatalf("Dump result to file err, Err: %v", err)
	}
	return err
}

func DumpToFile(filename string) (err error) {
	file, err := os.Create(filename)
	if err == nil {
		_, items := CacheStatus()
		for _, v := range items {
			result := v.Object.(ScanResult)
			file.WriteString(fmt.Sprintf("%v:%v|%v,%v:%v\n", result.Service.Ip, result.Service.Port,
				result.Service.Protocol, result.Service.Username, result.Service.Password))
		}
	}
	return err
}
