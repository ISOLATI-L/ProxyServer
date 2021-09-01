package cache

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"unsafe"
)

type Cache struct {
	Id         string
	Header     http.Header
	StatusCode int
	Body       []byte
}

func GetAbstract(header http.Header) string {
	headerData := *(*[]byte)(unsafe.Pointer(&header))
	abstract := fmt.Sprintf("%x", md5.Sum(headerData))
	return abstract
}

func Check(r *http.Request) (bool, error) {
	return false, nil
}
