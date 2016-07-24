package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func crawl(url string) {

	//resp, err := http.Get(url);
	//fmt.Println(resp.Body)
	//fmt.Println(err)

	//if err != nil {
	//	fmt.Printf("http get response errror=%s\n", err)
	//}

	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("body=%s\n", body)

	doc, err := goquery.NewDocument("http://sports.sina.com.cn")

	fmt.Println(doc)
	dhead := doc.Find("head")
	dcharset := dhead.Find("meta[http-equiv]")
	charset, _ := dcharset.Attr("content")

	fmt.Println(dhead.Html())
	fmt.Println(charset)
	fmt.Println(err)
}

func main() {
	crawl("http://www.baidu.com")
}
