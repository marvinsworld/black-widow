package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func crawl(url string) {

	resp, err := http.Get(url);
	//fmt.Println(resp.Body)
	//fmt.Println(err)

	if err != nil {
		fmt.Printf("http get response errror=%s\n", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("body=%s\n", body)
}

func main() {
	crawl("http://www.baidu.com")
}
