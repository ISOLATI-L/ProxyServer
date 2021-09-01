package cache

import (
	"ProxyServer/db"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"unsafe"
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

func Get(r *http.Request) *Cache {
	abstract, err := GetAbstract(r.Header)
	if err != nil {
		return nil
	}
	log.Println(hex.EncodeToString(abstract[:]))
	row := db.DB.QueryRow(
		`SELECT file FROM cache
		WHERE Cid=UNHEX(?)`,
		hex.EncodeToString(abstract[:]),
	)
	var cacheName string
	err = row.Scan(
		&cacheName,
	)
	if err != nil {
		return nil
	}
	cacheFile, err := os.OpenFile(
		cacheName,
		os.O_RDONLY,
		0666,
	)
	if err != nil {
		return nil
	}
	defer cacheFile.Close()
	cacheByte, err := io.ReadAll(cacheFile)
	if err != nil {
		return nil
	}
	cacheData := *(**Cache)(unsafe.Pointer(&cacheByte))
	return cacheData
}

func Save(cache *Cache) {
}
