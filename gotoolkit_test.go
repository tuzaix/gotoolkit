package gotoolkit_test

import (
	"fmt"
	. "gotoolkit"
	"testing"
)

func TestRedisKit(t *testing.T) {
	rdbPool := NewRedisPool("xxxxxx", 26390, 10, 30000, 500, 3000, 3000)
	conn := rdbPool.Get()
	defer conn.Close()
	fmt.Println(conn.Do("EXISTS", "haha"))
}

func TestDBKit(t *testing.T) {
	dbPool := NewDBPool("xxxxxx", "xxxxx", "xxxx", 3329, "doraemon", 2, 2)

	rows, err := dbPool.DB.Query("show tables")
	defer rows.Close()

	if err == nil {
		for rows.Next() {
			var table string
			rows.Scan(&table)
			fmt.Println(table)
		}
	}
}
