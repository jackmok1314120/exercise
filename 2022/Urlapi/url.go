package main

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	//"io/ioutil"
	"log"
	"math/big"
	"net/http"
)

func main() {

	http.HandleFunc("/login1", login1)
	http.HandleFunc("/login2", login2)
	http.HandleFunc("/login3", login3)
	http.HandleFunc("/savepub", SavePub)
	http.ListenAndServe("0.0.0.0:8080", nil)
}

type Resp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type Auth struct {
	Username string `json:"username"`
	Pwd      string `json:"password"`
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

type pub struct {
	Code string `json:"code"`
}

func SavePub(writer http.ResponseWriter, request *http.Request) {
	var pub pub
	if err := json.NewDecoder(request.Body).Decode(&pub); err != nil {
		request.Body.Close()
		log.Fatal(err)
	}
	fmt.Println(pub)
	src := []byte("admin")
	r := []byte("66270032207836837310581553683616995730865465858766525000504815570891615338765")
	s := []byte("85582170485247458343460474223710259150711674847154007837989376111743039678112")
	publ := pub.Code
	err := EccVerifySig(publ, src, r, s)
	if err != nil {
		fmt.Println("失败")
	}
	fmt.Println("成功", err)
}

//post接口接收json数据
func login1(writer http.ResponseWriter, request *http.Request) {
	var auth Auth
	if err := json.NewDecoder(request.Body).Decode(&auth); err != nil {
		request.Body.Close()
		log.Fatal(err)
	}
	fmt.Println(auth)
	var respon = "ok"
	if err := json.NewEncoder(writer).Encode(respon); err != nil {
		log.Fatal(err)
	}
}
func login3(writer http.ResponseWriter, request *http.Request) {
	var data data
	if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
		request.Body.Close()
		log.Fatal(err)
	}
	fmt.Println(data)
	var respon = "ok"
	if err := json.NewEncoder(writer).Encode(respon); err != nil {
		log.Fatal(err)
	}
}

//接收x-www-form-urlencoded类型的post请求或者普通get请求
func login2(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	username, uError := request.Form["username"]
	pwd, pError := request.Form["password"]

	var result Resp
	if !uError || !pError {
		result.Code = "401"
		result.Msg = "登录失败"
	} else if username[0] == "admin" && pwd[0] == "123456" {
		result.Code = "200"
		result.Msg = "登录成功"
		fmt.Println(request)
	} else {
		result.Code = "203"
		result.Msg = "账户名或密码错误"
		fmt.Println(request)
	}
	if err := json.NewEncoder(writer).Encode(result); err != nil {
		log.Fatal(err)
		//fmt.Println(err)
	}
}

type Signature struct {
	r *big.Int
	s *big.Int
}

////验证签名
//func Verify(c *handler.Context) error {
//	fmt.Println("start")
//	type Signature struct {
//		//R    []byte `json:"r" query:"r"`
//		//S    []byte `json:"s" query:"s"`
//		//Src  []byte `json:"src" query:"src"`
//		R    string `json:"r" query:"r"`
//		S    string `json:"s" query:"s"`
//		Src  string `json:"src" query:"src"`
//		Test string `json:"test" query:"test"`
//	}
//
//	sig := new(Signature)
//	fmt.Println("sig=", sig)
//	fmt.Println("src=", string(sig.Src))
//
//	//err := c.Binder(sig)
//	err := c.DefaultBinder(sig)
//	fmt.Println("sig=", sig)
//
//	fmt.Println("sig.s=", sig.S)
//	fmt.Println("sig.r=", sig.R)
//	fmt.Println("test=", sig.Test)
//	src := []byte(sig.Src)
//	r := []byte(sig.R)
//	s := []byte(sig.S)
//	fmt.Println("sig.Src=", src)
//	fmt.Println("sig.r=", r)
//	fmt.Println("sig.s=", s)
//	if err != nil {
//		fmt.Println("没进来")
//		return handler.NewError(c, err.Error())
//	}
//	err = library.EccVerifySig("util/library/EccPrivateKey.pem", src, r, s)
//	if err != nil {
//		fmt.Println("失败")
//		return handler.NewError(c, err.Error())
//	}
//	fmt.Println("end")
//	res := handler.NewSuccess()
//	return res.ResultOk(c)
//
//}

func EccVerifySig(dertext string, src []byte, rText, sText []byte) error {
	//读取公钥解码
	//fp, err := ioutil.ReadFile(filename)
	////fp, err := filename
	//if err != nil {
	//	return err
	//}
	////pem decode
	derteext := []byte(dertext)
	block, _ := pem.Decode(derteext)

	//解码der得到私钥
	derText := block.Bytes
	publicKeyInterface, err := x509.ParsePKIXPublicKey(derText)
	if err != nil {
		return err
	}
	publicKey, ok := publicKeyInterface.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("断言失败，非ecdsa公钥!\n")
	}

	//对原文生成哈希
	hash := sha256.Sum256(src)
	var r, s big.Int
	err = r.UnmarshalText(rText)
	if err != nil {
		return err
	}
	err = s.UnmarshalText(sText)
	if err != nil {
		return err
	}
	//使用公钥验证
	isok := ecdsa.Verify(publicKey, hash[:], &r, &s)

	if !isok {
		return errors.New("校验失败")
	}
	return nil
}
