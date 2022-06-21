package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
)

var (
	Dbhostsip  = "127.0.0.1:3306"
	Dbusername = "test01"
	Dbpassowrd = "12345678"
	Dbname     = "test"
)

type mysql_db struct {
	db *sql.DB //定义结构体
}

//Transaction返回信息处理
type Respbody struct {
	Hash        string `json:"hash"`
	Ver         int    `json:"ver"`
	VinSz       int    `json:"vin_sz"`
	VoutSz      int    `json:"vout_sz"`
	Size        int    `json:"size"`
	Weight      int    `json:"weight"`
	Fee         int    `json:"fee"`
	RelayedBy   string `json:"relayed_by"`
	LockTime    int    `json:"lock_time"`
	TxIndex     int64  `json:"tx_index"`
	DoubleSpend bool   `json:"double_spend"`
	Time        int    `json:"time"`
	BlockIndex  int    `json:"block_index"`
	BlockHeight int    `json:"block_height"`
	Inputs      []struct {
		Sequence int64  `json:"sequence"`
		Witness  string `json:"witness"`
		Script   string `json:"script"`
		Index    int    `json:"index"`
		PrevOut  struct {
			Spent             bool   `json:"spent"`
			Script            string `json:"script"`
			SpendingOutpoints []struct {
				TxIndex int64 `json:"tx_index"`
				N       int   `json:"n"`
			} `json:"spending_outpoints"`
			TxIndex int64  `json:"tx_index"`
			Value   int    `json:"value"`
			Addr    string `json:"addr"`
			N       int    `json:"n"`
			Type    int    `json:"type"`
		} `json:"prev_out"`
	} `json:"inputs"`
	Out []struct {
		Type              int  `json:"type"`
		Spent             bool `json:"spent"`
		Value             int  `json:"value"`
		SpendingOutpoints []struct {
			TxIndex int64 `json:"tx_index"`
			N       int   `json:"n"`
		} `json:"spending_outpoints"`
		N       int    `json:"n"`
		TxIndex int64  `json:"tx_index"`
		Script  string `json:"script"`
		Addr    string `json:"addr"`
	} `json:"out"`
}
type temp struct {
	Coin    string
	tx_hash []struct {
		tx_hash string
	}
}

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func CheckTxHash(url string) (string, int, error) {

	client := &http.Client{}
	// get请求
	req, err := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	if err != nil {
		return "1", 0, err
	}
	data, err := ioutil.ReadAll(resp.Body)

	//fmt.Println("response:", string(data))
	var resul Respbody
	if err := json.Unmarshal(data, &resul); err != nil {
		return " 2", 0, nil
	}
	for i := 0; i < len(resul.Inputs); i++ {
		//fmt.Println("Transaction:", resul.Hash, "\n", "输出：", i+1, resul.Inputs[i].PrevOut.Addr, resul.Inputs[i].PrevOut.Value, "\t")
		input := resul.Inputs[i]
		Iaddr := input.PrevOut.Addr
		Ivalue := input.PrevOut.Value
		return Iaddr, Ivalue, err
	}
	for i := 0; i < len(resul.Out); i++ {
		//fmt.Println("输入：", i+1, resul.Out[i].Addr, "\t")
		output := resul.Out[i]
		Oaddr := output.Addr
		Ovalue := output.Value
		return Oaddr, Ovalue, err
		//return resul.Out[i].Addr, resul.Out[i].Value, nil
	}

	//return resul.Out[0].Addr + "\t" + resul.Out[0].Addr, resul.Out[0].Value, nil
	return "", 0, nil
}

func (f *mysql_db) mysql_open() { //打开
	db, err := sql.Open("mysql", Dbusername+":"+Dbpassowrd+"@tcp("+Dbhostsip+")/"+Dbname)
	if err != nil {
		fmt.Println("链接失败")
	}
	//fmt.Println("链接数据库成功...........已经打开")
	f.db = db
}
func (f *mysql_db) mysql_close() { //关闭
	defer f.db.Close()
	fmt.Println("数据库...........已经关闭")
}

/*func (f *mysql_db) mysql_select(sqldata string) string {

	result := make([]temp, 0)

	rows, err := f.db.Query(sqldata)
	if err != nil {
		println(err)
	}
	for rows.Next() {
		var t temp
		err = rows.Scan(&t.tx_hash)
		if err != nil {
			panic(err)
		}

		result = append(result, t)
	}
	return result

}*/
func (f *mysql_db) mysql_select(sql_data string) {
	rows, err := f.db.Query(sql_data)
	if err != nil {
		println(err)
	}
	var tx_hash string
	for rows.Next() {

		err = rows.Scan(&tx_hash)
		if err != nil {
			println(err)
		}

		//url := "https://blockchain.info/rawtx/" + tx_hash
		//CheckTxHash(url)
		//fmt.Println(CheckTxHash(url)) //打印资料
		//取出tx_hash
		//打印tx_hash
		//fmt.Println(tx_hash)

	}

}

func main() {
	var db mysql_db
	sqld := "SELECT tx_hash FROM temp"
	db.mysql_open()
	db.mysql_select(sqld)
	//fmt.Println("tx:", tx_hash)
	db.mysql_close() //关闭
	//url := "https://blockchain.info/rawtx/" + tx_hash
	//CheckTxHash(url)
}
