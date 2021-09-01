package cache

import (
	"crypto/md5"
	"encoding/json"
	"net/http"
)

type Cache struct {
	Id         string
	Header     http.Header
	StatusCode int
	Body       []byte
}

// 获取请求的md5校验值
func GetAbstract(header http.Header) ([16]byte, error) {
	headerData, err := json.Marshal(header)
	if err != nil {
		return [16]byte{}, err
	}
	abstract := md5.Sum(headerData)
	return abstract, nil
}

func Check(r *http.Request) (bool, error) {
	return false, nil
}
