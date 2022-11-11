package main

import (
	"log"
	"net/http"

	"github.com/lestrrat-go/libxml2"
	"github.com/lestrrat-go/libxml2/xpath"
	"github.com/pkg/errors"
)

var baseUrl = "http://smse.whut.edu.cn/yjspy/xsdw/bdxx/"
var userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36"
var client = &http.Client{}

func main() {
	body, err := getReqBody(baseUrl, "GET")
	if err != nil {
		log.Printf("%+v\n", err)
	}
	log.Printf("%+v\n", body)
}

func getReqBody(url string, method string) (*http.Response, error) {
	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return nil, errors.Wrap(err, "创建httpRequest错误")
	}
	req.Header.Add("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "请求错误")
	}
	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "读取body错误")
	// }
	
	return resp, nil
}
func xpathDoc(url string, method string) error {
	reqBody, err := getReqBody(url, method)
	if err != nil {
		return err
	}
	defer reqBody.Body.Close()
	doc, err := libxml2.ParseHTMLReader(reqBody.Body)
	if err != nil {
		return errors.Wrap(err, "创建ParseHTMLReader失败")
	}
	defer doc.Free()
	nodes := xpath.NodeList(doc.Find(`//ul[@class="normal_list2"]/li`))
	log.Println("11111", len(nodes))
	_ = nodes
	return nil
}
