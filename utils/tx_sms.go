package utils

import (
	"bytes"
	"crypto/tls"
	"github.com/json-iterator/go"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strconv"
	"time"
)

const (
	verifyCodeUrl = "https://yun.tim.qq.com/v5/tlssmssvr/sendsms?"
	appid         = "1400047383"
	appKey        = "41e5490bf9fb98c68771c73586189780"
	tpl_id        = 51653
	sign          = "开卷小记"
)

/*
{
    "tel": { //如需使用国际电话号码通用格式，如："+8613788888888" ，请使用sendisms接口见下注
        "nationcode": "86", //国家码
        "mobile": "13788888888" //手机号码
    },
    "sign": "腾讯云", //短信签名，如果使用默认签名，该字段可缺省
    "tpl_id": 19, //业务在控制台审核通过的模板ID
     //假定这个模板为：您的{1}是{2}，请于{3}分钟内填写。如非本人操作，请忽略本短信。
    "params": [
        "验证码",
        "1234",
        "4"
    ], //参数，分别对应上面假定模板的{1}，{2}，{3}
    "sig": "ecab4881ee80ad3d76bb1da68387428ca752eb885e52621a3129dcf4d9bc4fd4", //app凭证，具体计算方式见下注
    "time": 1457336869, //unix时间戳，请求发起时间，如果和系统时间相差超过10分钟则会返回失败
    "extend": "", //通道扩展码，可选字段，默认没有开通(需要填空)。
    //在短信回复场景中，腾讯server会原样返回，开发者可依此区分是哪种类型的回复
    "ext": "" //用户的session内容，腾讯server回包中会原样返回，可选字段，不需要就填空。

	"sig"字段根据公式sha256(appkey=$appkey&random=$random&time=$time&mobile=$mobile)生成
}

{
    "result": 0, //0表示成功(计费依据)，非0表示失败
    "errmsg": "OK", //result非0时的具体错误信息
    "ext": "", //用户的session内容，腾讯server回包中会原样返回
    "sid": "xxxxxxx", //标识本次发送id，标识一次短信下发记录
    "fee": 1 //短信计费的条数
}
*/

func SMS_SendVerifyCode(phone string) (result *TX_SMS_CodeInfo, code string, err error) {
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}
	client := http.Client{Transport: tr}
	client.Jar, _ = cookiejar.New(nil)

	var curTimeInt = time.Now().Unix()
	var curTime = strconv.FormatInt(curTimeInt, 10)
	var random = bson.NewObjectId().Hex()

	var verifyCode = strconv.Itoa(RandInt(1001, 9999))

	url := verifyCodeUrl + "sdkappid=" + appid + "&random=" + random

	sigStr := "appkey=" + appKey + "&random=" + random + "&time=" + curTime + "&mobile=" + phone
	sig := GetSha256Code(sigStr)

	body := map[string]interface{}{
		"tel": map[string]interface{}{
			"nationcode": "86",
			"mobile":     phone,
		},
		"tpl_id": tpl_id,
		"params": []string{verifyCode, "30"},
		"sig":    sig,
		"time":   curTimeInt,
	}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	bodyData, _ := json.Marshal(&body)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyData))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
		return nil, verifyCode, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
		return nil, verifyCode, err
	}

	var codeInfo TX_SMS_CodeInfo
	err = json.Unmarshal(b, &codeInfo)
	if err != nil {
		panic(err)
		return nil, verifyCode, err
	}

	return &codeInfo, verifyCode, nil
}

type TX_SMS_CodeInfo struct {
	Result int    `json:"result"`
	ErrMsg string `json:"errmsg"`
	Ext    string `json:"ext"`
	Sid    string `json:"sid"`
	Fee    int    `json:"fee"`
}
