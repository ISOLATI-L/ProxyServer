package cache

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
)

type Cache struct {
	Id         string
	Header     http.Header
	StatusCode int
	Body       []byte
}

func GetAbstract(header http.Header) (string, error) {
	headerData, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	abstract := fmt.Sprintf("%x", md5.Sum(headerData))
	return abstract, nil
}

func Check(r *http.Request) (bool, error) {
	return false, nil
}
