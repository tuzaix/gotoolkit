package gotoolkit

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type DBPool struct {
	*sql.DB
}

/*
   使用默认的charset utf8
*/
func NewDBPool(username string, password string, host string, port int, database string, maxIdle int, maxOpen int) *DBPool {
	return NewDBPoolWithCharset(username, password, host, port, database, maxIdle, maxOpen, "utf8")
}

/*
	指定字符集的mysql链接
*/
func NewDBPoolWithCharset(username string, password string, host string, port int, database string, maxIdle int, maxOpen int, charset string) *DBPool {
	var (
		url  = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", username, password, host, port, database, charset)
		conn *sql.DB
		err  error
	)
	if conn, err = sql.Open("mysql", url); err != nil {
		panic("建立链接失败：" + url + "\t" + err.Error())
	}

	if err = conn.Ping(); err != nil {
		panic("Ping链接失败：" + err.Error())
	}
	conn.SetMaxIdleConns(maxIdle)
	conn.SetMaxOpenConns(maxOpen)
	return &DBPool{conn}
}
