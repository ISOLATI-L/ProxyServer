package blacklist

import (
	"proxy/db"
)

func Check(host string) (bool, error) {
	row := db.DB.QueryRow(
		`SELECT COUNT(host) FROM blacklist
		WHERE host=? || ip=?;`,
		host,
		host,
	)
	var count int
	err := row.Scan(
		&count,
	)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
