
package util

import (
	"tt-crack/logger"
	"tt-crack/vars"

	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadDbPort(fileName string)(DbPort map[string]int,err error){
	dbPortFile, err := os.Open(fileName)
	if err != nil {
		logger.Log.Fatalf("Open ip List file err, %v", err)
	}

	defer dbPortFile.Close()

	scanner := bufio.NewScanner(dbPortFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan(){
		line :=scanner.Text()
		if line ==""{
			continue
		}
		dbPort := strings.TrimSpace(line)
		t:= strings.Split(dbPort,":")
		//portname := t[0]
		portint, err := strconv.Atoi(t[1])
		if err == nil{
			vars.DbPort[t[0]]=portint

		}
	}
	return vars.DbPort,err
}


func ReadUserDict(userDict string) (users []string, err error) {
	file, err := os.Open(userDict)
	if err != nil {
		logger.Log.Fatalf("Open user dict file err, %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		user := strings.TrimSpace(scanner.Text())
		if user != "" {
			users = append(users, user)
		}
	}
	return users, err
}

func ReadPasswordDict(passDict string) (password []string, err error) {
	file, err := os.Open(passDict)
	if err != nil {
		logger.Log.Fatalf("Open password dict file err, %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		passwd := strings.TrimSpace(scanner.Text())
		if passwd != "" {
			password = append(password, passwd)
		}
	}
	password = append(password, "")
	return password, err
}
