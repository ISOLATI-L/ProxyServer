package cache

import (
	"ProxyServer/db"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
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

// 获取缓存
func Get(abstract [16]byte) *Cache {
	// abstract, err := GetAbstract(r.Header)
	// if err != nil {
	// 	return nil
	// }
	log.Println(hex.EncodeToString(abstract[:]))
	row := db.DB.QueryRow(
		`SELECT file FROM cache
		WHERE Cid=UNHEX(?)`,
		hex.EncodeToString(abstract[:]),
	)
	var cacheName string
	err := row.Scan(
		&cacheName,
	)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println("Error: ", err.Error())
		}
		return nil
	}
	cacheFile, err := os.OpenFile(
		cacheName,
		os.O_RDONLY,
		0666,
	)
	if err != nil {
		log.Println("Error: ", err.Error())
		return nil
	}
	defer cacheFile.Close()
	cacheByte, err := io.ReadAll(cacheFile)
	if err != nil {
		log.Println("Error: ", err.Error())
		return nil
	}
	cacheData := *(**Cache)(unsafe.Pointer(&cacheByte))
	return cacheData
}

// 保存缓存
func Save(abstract [16]byte, cache *Cache) {
	cacheName := "CacheFiles/" + hex.EncodeToString(abstract[:])
	cacheFile, err := os.OpenFile(
		cacheName,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}
	cacheByte := *(*[]byte)(unsafe.Pointer(cache))
	_, err = cacheFile.Write(cacheByte)
	if err != nil {
		log.Println("Error: ", err.Error())
		cacheFile.Close()
		return
	}
	err = cacheFile.Close()
	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	result, err := db.DB.Exec(
		`INSERT INTO cache (Cid, file)
		VALUES (UNHEX(?), ?);`,
		hex.EncodeToString(abstract[:]),
		cacheName,
	)
	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}
	if affected == 0 {
		log.Println("Error: ", errors.New("affected 0 rows"))
		return
	}
}
