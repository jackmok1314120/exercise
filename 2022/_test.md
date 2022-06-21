# API 简介

欢迎使用 虎符托管API！ 你可以使用此 API 获得账单数据，进行提现，并且管理你的资产。

# 接口

以下是虎符托管API提供的API接口：

接口说明:

|  请求方式  |               接口               |      介绍       |
|  -------  | ------------------------        | -------------- |
|<b> POST</b>      | {IPAddress}/api/open/v1/withdraw           | 商户发起提现     |
|<b> GET</b>       | {IPAddress}/api/open/v1/searchBill         | 商户账单查询     |
|<b> GET</b>       | {IPAddress}/api/open/vip/v1/findAssets     | 币种资产查询接口  |
|<b> GET</b>       | {IPAddress}/api/open/vip/v1/singleAddress  | 单个拉取地址     |
|<b> POST</b>      | {IPAddress}/api/open/vip/v1/batchAddress   | 批量拉取地址     |

`{IPAddress}`说明：该参数为虎符提供的端口地址或者域名。

正式环境：`manager.api.hoocustody.com`

如：
`http://manager.api.hoocustody.com`

拼接后

     `http://manager.api.hoocustody.com/api/open/v1/withdraw`

## 接入说明

接口仅支持<b>`GET`</b>和<b>`POST`</b>请求

1) 所有请求参数按照按字符顺序进行排序后，以「key=value&key1=value」的方式组合成字符串然
   后进行加密，将加密后的十六进制字符串放入参数sign中一并传入

2) sign加密方式为：HMAC-SHA256，测试key：12345，加密结果以十六进制表示，参考golang示例

3) 接口统一返回格式：

> 统一返回格式：

```shell
   {
     "code": "10000",
     "message": "success",
     "data": {}
   }

```

code参数说明：

     code=10000时接口成功，其他code代码为接口请求错误,具体错误信息可以参考 错误码
    「message」为请求附带的提示信息
    「data」为返回的具体数据

## 商户发起提现

HTTP请求

<b>POST</b> `{IPAddress}/api/open/v1/withdraw`

> Request:

```shell
{
  "fromAddress": "16Tz7F429szTnjT3yvKwsr6n1fhG9w4woD",
  "toAddress": "18h1VbLBmxoeL44gnWxPC4NE8HampsqoDt",
  "sign": "111",
  "coin": "BTC",
  "chain": "BTC",
  "clientId": "XbEIfaFkrPeDCBDdwHOFujHdNPrJPCFX",
  "coinNums": 0.0011,
  "memo": "BTC"
}

```

请求参数:

|参数名|参数类型|是否必须|描述|
|-----|-------|--------|---|
| fromAddress |   string    |  是   |  提现地址 |
| toAddress |    string  |  是   |    接收地址   |
| sign |   string  |    是    |   签名 |
| coin |   string  |    是    |   币种名称，如BTC、ETH等 |
| chain |   string  |    是    |   主链币名称，如BTC、ETH等 |
| clientId |  string  |    是    |   商户的业务线ClintId |
| coinNums |   string/float  |    是    |   提现金额>0 |
| memo |   string  |    否    |   备注 |

**备注**

    虎符托管不关心Memo内容，全量推送交易信息

>sign 加密「key=value」排列：

```shell
fromAddress=value&toAddress=value&coin=value&coinNums=value&memo=value&clientId=value
```

> Responds:

```shell
{
  "message": "success",
  "code": "10000",
  "data": {
    "serialNo": "账单号"
  }
}

```

基本参数：

|参数名|参数类型 |描述|
|-----|-------|--------|
|message |string |返回提示信息|
|code |string |10000为成功，其他为失败|
|data |object |返回的具体数据|

data内容：

|参数名|参数类型 |描述|
|-----|-------|--------|
| serialNo |   string |  账单号 |

## 商户账单查询

HTTP请求

<b>GET</b> `{IPAddress}/api/open/v1/searchBill`

> Request:

```shell
{IPAddress}/api/open/v1/searchBill?serialNo=XXX&clientId=XXX&sign=XXX
```

请求参数:

|参数名|参数类型|    是否必须|描述|
|-----|-------|--------|---|
| serialNo |   string    |  是   |  账单号 |
| clientId |    string  |  是   |    商户的业务线ClintId   |
| sign |   string  |    是    |   签名 |
> sign 加密「key=value」排列：

```shell
serialNo=value&clientId=value
```

> Responds:

```shell
{
  "code": 10000,
  "data": {
    "auditTime": "",
    "billStatusName": "充值确认中",
    "chainName": "",
    "coinName": "",
    "confirmTime": "",
    "destroyFee": "0",
    "fee": "0",
    "id": 0,
    "memo": "",
    "phone": "",
    "realNums": "0",
    "remark": "",
    "resultName": "待审核",
    "serialNo": "",
    "serviceName": "",
    "txFromAddr": "",
    "txId": "",
    "txTime": "",
    "txToAddr": "",
    "txTypeName": "链上地址转入-正在接收",
    "upChainFee": "0"
  },
  "message": "success"
}
```

基本参数：

|参数名|参数类型 |描述|
|-----|-------|--------|
|message |string |返回提示信息|
|code |string |10000为成功，其他为失败|
|data |object |返回的具体数据|

data内容：

|参数名|参数类型 |描述|
|-----|-------|--------|
| id |   int |  id |
| txId |   string |  交易ID |
| serialNo |   string |  账单号 |
| chainName |   string |  主链币名称 |
| coinName |   string |  代币名称 |
| fee |   string |  提现手续费 |
| destroyFee |   string |  销毁数量 |
| upChainFee |   string |  矿工费 |
| phone |   string |  手机号 |
| nums |   string |  提现金额 |
| realNums |   string |  实际到账 |
| serviceName |   string |  业务线名称 |
| txTypeName |   string |  交易类型 |
| billStatusName |   string |  账单状态 |
| resultName |   string |  审核状态 |
| txToAddr |   string |  接收地址 |
| txFromAddr |   string |  发送地址 |
| txTime |   string |  交易时间 |
| auditTime |   string |  审核时间 |
| confirmTime |   string |  确认时间 |
| memo |   string |  memo |
| remark |   string |  审核备注 |

## 币种资产查询接口

HTTP请求

<b>GET</b> `{IPAddress}/api/open/v1/findAssets`

> Request:
```shell
{IPAddress}/api/open/v1/findAssets?chain=XXX&coin=XXX&clientId=XXX&sign=XXX&limit=10&
offset=0
```

请求参数:

|参数名|参数类型|    是否必须|描述|
|-----|-------|--------|---|
| chain |   string    |  是   |  主链币名 |
| coin |   string    |  是   |  代币名 |
| clientId |    string  |  是   |    商户的业务线ClintId   |
| sign |   string  |    是    |   签名 |
| limit |   int    |  是   |  查询数量 |
| offset |   int    |  是   |  起始位置从0开始 |

> sign 加密「key=value」排列：

```shell
chain=value&coin=value&clientId=value&limit=value&offset=value
```
> Responds:

```shell
{
  "code": 10000,
  "data": {
    "list": [
      {
        "chainName": "BTC",
        "coinName": "BTC",
        "coinPrice": "44129.74",
        "nums": "100",
        "reduced": "27978255.16",
        "valuation": "4412974"
      }
    ],
    "total": 1
  },
  "message": "success"
}
```

基本参数：
基本参数：

|参数名|参数类型 |描述|
|-----|-------|--------|
|message |string |返回提示信息|
|code |string |10000为成功，其他为失败|
|data |object |返回的具体数据|

data内容：

|参数名|参数类型 |描述|
|-----|-------|--------|
| chainName |   int |  主链名 |
| coinName |   string |  代币名 |
| coinPrice |   string |  代币价格（usdt） |
| nums |   string |  资产数量 |
| reduced |   string |  折算RMB |
| valuation |   string |  估值 |

**备注**

    reduced默认折算为RMB，估值计算为：nums*coinPrice，即为数量乘当时的代币交易所币价；
    具体提现金额以提现交易时市场价为准，该coinPrice所显示价格仅供参考，不作为保证。
## 拉取单个地址

HTTP请求

<b>GET</b> ` {IPAddress}/api/open/v1/singleAddress`

> Request:

```shell
{IPAddress}/api/open/v1/singleAddress?chain=XXX&coin=XXX&clientId=XXX&sign=XXX&userId=XXX
```

请求参数:

|参数名|参数类型| 是否必须 |描述|
|-----|-------|------|---|
| chain |   string    | 是    |  主链币名 |
| coin |   string    | 是    |  代币名 |
| clientId |    string  | 是    |  商户的业务线ClintId   |
| sign |   string  | 是    |   签名 |
| userId |   string    | 是    |  地址的标识，用于区分/绑定地址 |

> sign 加密「key=value」排列：

```shell
chain=value&coin=value&clientId=value&userId=value
```

> Responds:

```shell
{
  "code": 10000,
  "data": {
    "address": "地址"
  },
  "message": "success"
}
```

基本参数：

|参数名|参数类型 |描述|
|-----|-------|--------|
|message |string |返回提示信息|
|code |string |10000为成功，其他为失败|
|data |object |返回的具体数据|

data内容：

|参数名|参数类型 |描述|
|-----|-------|--------|
| address |   string |  地址 |

## 批量拉取地址

HTTP请求

<b>POST</b> ` {IPAddress}/api/open/v1/batchAddress`

> Request:

```shell
{
    "userId":["1","2","3","4"],
    "coin":"TRX",
    "chain":"TRX",
    "clientId":"xxxxx",
    "sign":"aaa",
    "nums":2
}
```

请求参数:

|参数名|参数类型| 是否必须 |描述|
|-----|-------|------|---|
| chain |   string    | 是    |  主链币名 |
| coin |   string    | 是    |  代币名 |
| nums |   int    | 是    |  拉取的地址数 |
| clientId |    string  | 是    |  商户的业务线ClintId   |
| sign |   string  | 是    |   签名 |
| userId |   []string    | 是    |  地址的标识数组，用于区分/绑定地址 |
> sign 加密「key=value」排列：

```shell
chain=value&coin=value&clientId=value&nums=value
```

> Responds:

```shell
{
  "code": 10000,
  "data": [
    {
      "userId": "",
      "address": ""
    }
  ],
  "message": "success"
}
```
基本参数：

|参数名|参数类型 |描述|
|-----|-------|--------|
|message |string |返回提示信息|
|code |string |10000为成功，其他为失败|
|data |object |返回的具体数据|

data内容：

|参数名|参数类型 |描述|
|-----|-------|--------|
| userId |   string |  商家填入的用户ID |
| address |   string |  地址 |


## 商家回调地址

商家提供回调地址，用于账单数据推送。

###请求方法：

<b>`POST`</b>
###接收参数：

shell格式，接收字段：`data`

```shell
{
  "data": {
    "auditTime": "",
    "billStatusName": "充值确认中",
    "chainName": "",
    "coinName": "",
    "confirmTime": "",
    "destroyFee": "0",
    "fee": "0",
    "id": 0,
    "memo": "",
    "nums": "0",
    "phone": "",
    "realNums": "0",
    "remark": "",
    "resultName": "待审核",
    "serialNo": "",
    "serviceName": "",
    "txFromAddr": "",
    "txId": "",
    "txTime": "",
    "txToAddr": "",
    "txTypeName": "链上地址转入-正在接收",
    "upChainFee": "0"
  }
}
```

data内容：

|参数名|参数类型|描述|
|-----|-------|--------|
| id |   int |  id |
| txId |   string |  交易ID |
| serialNo |   string |  账单号 |
| chainName |   string |  主链币名称 |
| coinName |   string |  代币名称 |
| fee |   string |  提现手续费 |
| destroyFee |   string |  销毁数量 |
| upChainFee |   string |  矿工费 |
| phone |   string |  手机号 |
| nums |   string |  提现金额 |
| realNums |   string |  实际到账 |
| serviceName |   string |  业务线名称 |
| txTypeName |   string |  交易类型 |
| billStatusName |   string |  账单状态 |
| resultName |   string |  审核状态 |
| txToAddr |   string |  接收地址 |
| txFromAddr |   string |  发送地址 |
| txTime |   string |  交易时间 |
| auditTime |   string |  审核时间 |
| confirmTime |   string |  确认时间 |
| memo |   string |  memo |
| remark |   string |  审核备注 |


## sign签名demo

> 以Go代码为例：

```go
package main
import (
        "crypto/hmac"
        "crypto/sha256"
        "encoding/hex"
        "fmt"
)
func main() {
        ToHmacSha256("key=value&key1=value1", "secretKey") // 调用HmacSha256加密
}
func ToHmacSha256(message string, secret string) string {
        key := []byte(secret)
        h := hmac.New(sha256.New, key)
        h.Write([]byte(message))
        fmt.Println(h.Sum(nil))
        sha := hex.EncodeToString(h.Sum(nil))
        fmt.Println(sha)
        return sha
}
```

> 以Python代码为例：

```python
import hmac
import hashlib
import requests
import shell

def hmac_sign(key, params):
    keys = sorted(params)
    sign_str = "&".join(["%s=%s" % (k, params.get(k)) for k in keys])
    print(sign_str)
    sign = hmac.new(key.encode("utf-8"), sign_str.encode("utf-8"), digestmod=hashlib.sha256).hexdigest()
    return sign_str, sign


if __name__ == "__main__":
    url = "https://***/api/open/vip/v1/accounts"
    secret_key = '12345'
    data = {
        "client_id": "testtest",
        "coin_name": "BTC",
  }
    sign_str, sign = hmac_sign(secret_key, data)
    data['sign'] = sign
    print(sign)

    # get请求
    # gUrl = url + "?" + sign_str
    # gRes = requests.get(gUrl, timeout=10)

    # post 请求
    res = requests.post(url, data=data, timeout=10)
    content = shell.loads(res.content)
    print(content)
    if content["code"] == "10000":
        print("success")
    else:
        print("failed")
```
###Go:

参数说明：

    「message」 加密排列好的字符串 key=value&key1=value1
    「secret」 secretKey，由托管生成业务线时生成secretKey和clientId

###Python:

参数说明：

    「message」 加密排列好的字符串 key=value&key1=value1
    「secret」 secretKey，由托管生成业务线时生成secretKey和clientId


## 错误码说明

| code | 说明 |
| ---- | --- |
|10101 | 用户不存在|
|10301 | 币种{coin}类型错误|
|10304 | 余额不足|
|10305 | 参数错误|
|10306 | 记录错误或不存在|
|10307 | 该地址与币种不匹配，请重新填写|
|10310 | 不能自己给自己发送|
|10321 | 低于最小发送限额|
|10322 | 超过最大发送限额|
|10325 | 地址存储失败|
|10326 | 生成地址失败|
|10327 | 提币异常|
|10328 | clientId 无效|
|10329 | 业务线 查询错误|
|10403 | 签名错误|
|10404 | 订单号不存在或重复|
|30006 | 币种错误|


