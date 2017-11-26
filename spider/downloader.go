package spider

import (
	"github.com/json-iterator/go"
	// "fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
)

const (
	Cookie string = `d_c0="ABCCZxtDnAuPTnewEeTL4U-MwxEDo5nZAsE=|1492245135"; _zap=20bbaae2-4038-4a13-80e9-069e57e766ed; q_c1=442a33f0784d4414b44622ff1a916311|1506603780000|1491708449000; q_c1=442a33f0784d4414b44622ff1a916311|1509522832000|1491708449000; l_cap_id="YTU2MGViZTA0MjcwNDM0YWExZWFiMWZkYmQ4Njk0ZDA=|1509590453|eb9e254e432ed77a663b0bd1e553283d6c6744be"; r_cap_id="OGYyYjZjZmRkYjVhNDRmMWFhZTg0MDQ2ZjVjYjlkYzM=|1509590453|42e5ad70edcae46c0da3e257583b3bd63bee1c8a"; cap_id="ZDliYzczYWIwZWNlNGZkMTlhNDFkZDc4NzY4YTEyYzQ=|1509590453|e2541157396cb9578c3b8e5d9f440d52d9273086"; z_c0="2|1:0|10:1509590516|4:z_c0|92:Mi4xZ2tSeEFBQUFBQUFBRUlKbkcwT2NDeWNBQUFDRUFsVk45QklpV2dDQmp3dVFzeVUtNWE2aUxNZDV1YmlMcUh4OVlB|93f2fa88072228f3a4ba094451f87049bd830028726a5aa5f2736ede926f1e03"; aliyungf_tc=AQAAAFNDhE3Y2wgABsrPb1/Ebg8C2N6z; s-t=autocomplete; s-q=golang%20%E7%88%AC%E8%99%AB; s-i=4; sid=s314r21g; __utma=51854390.1365864264.1511319719.1511319719.1511319719.1; __utmc=51854390; __utmz=51854390.1511319719.1.1.utmcsr=zhihu.com|utmccn=(referral)|utmcmd=referral|utmcct=/question/30285547; __utmv=51854390.100-1|2=registration_date=20140807=1^3=entry_date=20140807=1; _xsrf=5d962648-cc2f-4378-94d6-d7f7f8cb4bc2`
)

func GetJson(url string, obj interface{}) (body []byte, err error) {
	request, _ := http.NewRequest("GET", url, nil)
	_addDefaultHeader(request)

	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	// str := string(body[:])
	// fmt.Println(str)

	if obj != nil {
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		err = json.Unmarshal(body, obj)
	}
	return
}

func GetHtmlDoc(url string) (doc *goquery.Document, err error) {
	request, _ := http.NewRequest("GET", url, nil)
	_addDefaultHeader(request)

	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return
	}
	defer resp.Body.Close()

	doc, err = goquery.NewDocumentFromResponse(resp)

	return
}

func _addDefaultHeader(request *http.Request) {
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Set("Accept-Encoding", "")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7")
	request.Header.Set("Cookie", Cookie)
}
