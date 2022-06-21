// mysqldate
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xuri/excelize/v2"
	"log"
)

func main() {
	conn, err := sql.Open("mysql", "root:(HOO).com97@tcp(127.0.0.1:3306)/test?charset=utf8")
	defer conn.Close()
	res, err := conn.Exec("create table if not exists btc_tx(Coin varchar(20) default '' not null,Tx_hash varchar(40) default '' not null)")
	if err != nil {
		fmt.Println("exec1:", err)
	}
	log.Println(res, err)
	sql, err := conn.Prepare("insert into btc_tx(Cion ,Tx_hash ,) values (?,?)")
	log.Println(sql, err)
	defer sql.Close()

	f, err := excelize.OpenFile("/Users/jackmok/go/src/awesomeProject/2022/SlientMySql/BTC_tx.xlsx")
	if err != nil {
		fmt.Println("openfile", err)
		return
	}
	rows, err := f.GetRows("sheet1")
	//fmt.Println(len(rows[0]))
	//	fmt.Println(len(cell)
	count := 0
	for _, row := range rows {
		res, err := sql.Exec(row) //string格式
		if err != nil {
			log.Println(res, err)
		} else {
			count += 1
			log.Println("success")
		}
	}
	//fmt.Println(count)
	//fmt.Printf("%T", rows[4]) //string[]格式

}
