package hash

import (
	"tt-crack/vars"
)

func MakeTaskHash(k string) (string) {
	hash := MD5(k)
	return hash
}

func CheckTashHash(hash string) (bool) {
	_, ok := vars.SuccessHash[hash]
	return ok
}

func SetTaskHask(hash string) () {
	vars.Mutex.Lock()
	vars.SuccessHash[hash] = true
	vars.Mutex.Unlock()
}
