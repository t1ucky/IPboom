package util

import (
	"gopkg.in/cheggaaa/pb.v2"

	"tt-crack/logger"
	"tt-crack/models"
	"tt-crack/plugins"
	"tt-crack/util/hash"
	"tt-crack/vars"

	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

func GenerateTask(ipList []models.IpAddr, users []string, passwords []string) (tasks []models.Service, taskNum int) {
	tasks = make([]models.Service, 0)

	for _, user := range users {
		for _, password := range passwords {
			for _, addr := range ipList {
				service := models.Service{Ip: addr.Ip, Port: addr.Port, Protocol: addr.Protocol, Username: user, Password: password}
				tasks = append(tasks, service)
			}
		}
	}

	return tasks, len(tasks)
}

func RunTask(tasks []models.Service) {
	totalTask := len(tasks)
	vars.ProgressBar = pb.StartNew(totalTask)
	vars.ProgressBar.SetTemplate(`{{ rndcolor "Scanning progress: " }} {{  percent . "[%.02f%%]" "[?]"| rndcolor}} {{ counters . "[%s/%s]" "[%s/?]" | rndcolor}} {{ bar . "「" "-" (rnd "ᗧ" "◔" "◕" "◷" ) "•" "」" | rndcolor }} {{rtime . | rndcolor}} `)

	wg := &sync.WaitGroup{}

	// 创建一个buffer为vars.threadNum * 2的channel
	taskChan := make(chan models.Service, vars.ScanNum*2)

	// 创建vars.ThreadNum个协程
	for i := 0; i < vars.ScanNum; i++ {
		go crackPassword(taskChan, wg)
	}

	// 生产者，不断地往taskChan channel发送数据，直到channel阻塞
	for _, task := range tasks {
		wg.Add(1)
		taskChan <- task
	}

	close(taskChan)
	waitTimeout(wg, vars.TimeOut*2)
}

// 每个协程都从channel中读取数据后开始扫描并保存
func crackPassword(taskChan chan models.Service, wg *sync.WaitGroup) {
	for task := range taskChan {
		vars.ProgressBar.Increment()

		if vars.DebugMode {
			logger.Log.Debugf("checking: Ip: %v, Port: %v, [%v], UserName: %v, Password: %v, goroutineNum: %v", task.Ip, task.Port,
				task.Protocol, task.Username, task.Password, runtime.NumGoroutine())
		}

		var k string
		protocol := strings.ToUpper(task.Protocol)

		/*if protocol == "REDIS" || protocol == "FTP" || protocol == "SNMP" {
			k = fmt.Sprintf("%v-%v-%v", task.Ip, task.Port, task.Protocol)
		} else {*/
		k = fmt.Sprintf("%v-%v-%v", task.Ip, task.Port, task.Username)
		//}

		h := hash.MakeTaskHash(k)
		if hash.CheckTashHash(h) {
			wg.Done()
			continue
		}
		fn := plugins.ScanFuncMap[protocol]
		models.SaveResult(fn(task))
		wg.Done()
	}
}
var a string

// waitTimeout waits for the waitgroup for the specified max timeout.
// Returns true if waiting timed out.
func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}
