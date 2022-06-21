package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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

//表结构
type btc_transaction struct {
	CionNmae     string `db:"CionName"`
	Tx_hash      string `db:"Tx_hash"`
	OutPut_addr  string `db:"OutPut_addr"`
	Out_amount   int    `db:"Out_amount"`
	Input_addr   string `db:"Input_addr"`
	Input_amount int    `db:"Input_amount"`
}

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func (f *mysql_db) CheckTxHash(url string) (string, int, error) {
	//btc := btc_transaction{}
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
		input := resul.Inputs[i]
		Iaddr := input.PrevOut.Addr
		Ivalue := input.PrevOut.Value
		//return Iaddr, Ivalue, err

		//return resul.Inputs[i].PrevOut.Addr, resul.Inputs[i].PrevOut.Value, nil
		for i := 0; i < len(resul.Out); i++ {
			output := resul.Out[i]
			Oaddr := output.Addr
			Ovalue := output.Value
			//return Oaddr, Ovalue, err
			Inputins, err := f.db.Prepare("INSERT INTO btc_transactions (Tx_hash,Input_addr, Input_amount,Output_addr,Output_amount) VALUES(?,?,?,?,? )") // ?
			if err != nil {
				log.Printf("err：%v", err)
			}
			defer Inputins.Close()
			inp, err := Inputins.Exec(resul.Hash, Iaddr, Ivalue, Oaddr, Ovalue)
			if err != nil {
				fmt.Println("插入数据失败", err)
			} else {
				success, _ := inp.LastInsertId()
				insertId := strconv.Itoa(int(success))
				fmt.Println(strings.Join([]string{"成功插入数据：", insertId}, "\t-->\t"))
			}
		}
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
			panic(err)
		}

		url := "https://blockchain.info/rawtx/" + tx_hash
		fmt.Println(f.CheckTxHash(url))

	}

}

func main() {
	var db mysql_db
	sqld := "SELECT tx_hash FROM temp"
	db.mysql_open()
	db.mysql_select(sqld)
	//fmt.Println("tx:", Txhash)
	db.mysql_close() //关闭
	//url := "https://blockchain.info/rawtx/" + tx_hash
	//CheckTxHash(url)
}

/*func (f *mysql_db) mySql(contents []byte) {
	//用户名：密码^@tcp(地址:3306)/数据库
	db, err := sql.Open("mysql", Dbusername+":"+Dbpassowrd+"@tcp("+Dbhostsip+")/"+Dbname)
	if err != nil {
		fmt.Println(err)
		return
	}

	//查询表
	rows, err := db.Query("SELECT * FROM btc_transaction")

	//遍历打印
	for rows.Next() {
		var s btc_transaction
		err = rows.Scan(&s.CionNmae, &s.Tx_hash)
	}

	//执行MySql语句
	rg := regexp.MustCompile("")
	allSubmatch := rg.FindAllSubmatch(contents, -1)

	for _, m := range allSubmatch {
		//fmt.Printf("%s\n ", m[1])
		//fmt.Printf("%s\n ", m[2])
		//插入语句
		db.Exec("INSERT INTO btc_transaction(Tx_Hash,Output_addr,Out_amount,Input_addr,Input_amount)VALUES (?,?,?,?,?)", m[1], m[2])
	}
	rows.Close()
}*/
