package main

import (
	"database/sql"
	"fmt"
)

var (
	host = "127.0.0.1"
	port = "3306"
	name = "test01"
	//username = "test01"
	pwd      = "12345678"
	username = "test"
)

//var (
//	name     = "custody"
//	username = "custody"
//	pwd      = "OpDomFBx7Wk4Lts33qg57jXh"
//	host     = "3.115.25.173"
//	port     = "3306"
//)

type DB struct {
	db *sql.DB //定义结构体
}

func (D *DB) OpenDB() { //打开

	db, err := sql.Open(name, username+":"+pwd+"@tcp("+host+":"+port+")")
	if err != nil {
		fmt.Println("链接失败")
	}
	fmt.Println("链接数据库成功...........已经打开")
	D.db = db
}
func (D *DB) CloseDB() { //关闭
	defer D.db.Close()
	fmt.Println("数据库...........已经关闭")
}

type adminApi struct {
	id      string `db:"id"`
	name    string `db:"name"`
	path    string `db:"path"`
	method  string `db:"method"`
	tag     string `db:"tag"`
	created string `db:"created_at"`
	updated string `db:"updated_at"`
	deleted string `db:"deleted_at"`
}

func (D *DB) ReadDb(sqlData string) {
	rows, err := D.db.Query(sqlData)
	if err != nil {
		println(err)
	}
	//var tx_hash, Coin string

	for rows.Next() {
		var a adminApi //数据库内数据结构
		err = rows.Scan(&a.id, a.name)
		if err != nil {
			panic(err)
		}
		fmt.Sprintf(a.id, a.name)
		//将txhash传入btc查询url

	}

}
func main() {
	var d DB
	sqlDate := "SELECT id,name FROM admin_api"
	d.OpenDB()
	d.ReadDb(sqlDate)
}
