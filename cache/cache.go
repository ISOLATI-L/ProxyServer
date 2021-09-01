package cache

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
)

type Cache struct {
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

func GetCache(r *http.Request) *Cache {
	abstract, err := GetAbstract(r.Header)
	if err != nil {
		return nil
	}
	log.Println(hex.EncodeToString(abstract[:]))
	// db.DB.QueryRow(
	// 	`SELECT COUNT(Cid) FROM blacklist
	// 	WHERE Cid=UNHEX(?)`,
	// 	hex.EncodeToString(abstract[:]),
	// )
	return &Cache{}
}
