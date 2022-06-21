package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type auth struct {
	Username string `json:"username"`
	Pwd      string `json:"password"`
}

func main() {
	//get()
	postWithJson()
	//postWithUrlencoded()
}

func get() {
	//get请求
	//http.Get的参数必须是带http://协议头的完整url,不然请求结果为空
	resp, _ := http.Get("http://localhost:8080/login2?username=admin&password=123456")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	fmt.Printf("Get request result: %s\n", string(body))
}

type data struct {
	AuditTime      string `json:"auditTime"`
	BillStatusName string `json:"billStatusName"`
	ChainName      string `json:"chainName"`
	CoinName       string `json:"coinName"`
	ConfirmTime    string `json:"confirmTime"`
	DestroyFee     string `json:"destroyFee"`
	Fee            string `json:"fee"`
	Id             int    `json:"id"`
	Memo           string `json:"memo"`
	Nums           string `json:"nums"`
	Phone          string `json:"phone"`
	RealNums       string `json:"realNums"`
	Remark         string `json:"remark"`
	ResultName     string `json:"resultName"`
	SerialNo       string `json:"serialNo"`
	ServiceName    string `json:"serviceName"`
	TxFromAddr     string `json:"txFromAddr"`
	TxId           string `json:"txId"`
	TxTime         string `json:"txTime"`
	TxToAddr       string `json:"txToAddr"`
	TxTypeName     string `json:"txTypeName"`
	UpChainFee     string `json:"upChainFee"`
}

func postWithJson() {
	//post请求提交json数据
	auths := data{"1", "2", "3", "4", "5", "6", "7", 1, "8", "9", "123", "10", "unknow", "ing",
		"12344", "unknow", "fjujvndjv", "fghjkmnbgyu8987yh", "huio9", "fghjhbvftyui87tf", "", ""}
	ba, _ := json.Marshal(auths)
	resp, _ := http.Post("http://localhost:8080/login3", "application/json", bytes.NewBuffer([]byte(ba)))
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Post request with json result: %s\n", string(body))
}

/*func postWithUrlencoded() {
	//post请求提交application/x-www-form-urlencoded数据
	form := make(url.Values)
	form.Set("username", "admin")
	form.Add("password", "123456")
	resp, _ := http.PostForm("http://localhost:8080/login2", form)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Post request with application/x-www-form-urlencoded result: %s\n", string(body))
}*/
