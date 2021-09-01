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
)

type Cache struct {
	Header     http.Header
	StatusCode int
	Body       []byte
}

// 获取请求的md5校验值
func GetAbstract(r *http.Request) ([16]byte, error) {
	// log.Println(r.Host)
	// log.Println(r.RequestURI)

	headerData, err := json.Marshal(r.Header)
	if err != nil {
		return [16]byte{}, err
	}
	headerData = []byte(r.RequestURI + ":" + string(headerData))
	// log.Println(string(headerData))
	abstract := md5.Sum(headerData)

	// log.Println(hex.EncodeToString(abstract[:]))
	return abstract, nil
}

// 获取缓存
func Get(abstract [16]byte) *Cache {
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
	var cacheData *Cache
	err = json.Unmarshal(cacheByte, cacheData)
	if err != nil {
		log.Println("Error: ", err.Error())
		return nil
	}
	return cacheData
}

// 保存缓存
func Save(abstract [16]byte, cache *Cache) {
	_, err := os.Stat("CacheFiles")
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir("CacheFiles", 0777)
			if err != nil {
				log.Println("Error: ", err.Error())
				return
			}
		} else {
			log.Println("Error: ", err.Error())
			return
		}
	}

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

	cacheByte, err := json.Marshal(cache)
	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}
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
