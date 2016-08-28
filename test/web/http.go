package web

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tebeka/selenium"
	"github.com/sclevine/agouti"
	"log"
	"github.com/fedesog/webdriver"
	"time"
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

	doc, _ := goquery.NewDocument(url)

	//fmt.Println(doc)
	dbody := doc.Find("body")
	//dcharset := dhead.Find("meta[http-equiv]")
	//charset, _ := dcharset.Attr("content")

	//fmt.Println(dbody.Html())
	//fmt.Println(charset)
	//fmt.Println(err)

	dbody.Find("div.c-container").Each(func(i int, selection *goquery.Selection) {
		title, _ := selection.Html()
		fmt.Println(title)
	})
}

func selen() {
	var code = `
package main
import "fmt"

func main() {
	fmt.Println("Hello WebDriver!\n")
}
`

	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, _ := selenium.NewRemote(caps, "")
	defer wd.Quit()

	// Get simple playground interface
	wd.Get("http://play.golang.org/?simple=1")

	// Enter code in textarea
	elem, _ := wd.FindElement(selenium.ByCSSSelector, "#code")
	elem.Clear()
	elem.SendKeys(code)
}

func ago() {
	driver := agouti.Selenium() // 设置 driver
	driver.Start()
	page, _ := driver.NewPage(agouti.Browser("chrome"))

	page.Navigate("http://man.pay.test.dmall.com:8080/bestcard/list")

	fmt.Println(page.Find("header").Text())
	page.Find("#su").Click()
	////page.Click(SingleClick,LeftButton)
	fmt.Println(page.URL())
	//
	driver.Stop() // 关闭 driver
}

func fed() {
	chromeDriver := webdriver.NewChromeDriver("/usr/local/Cellar/chromedriver/2.22/bin/chromedriver")
	err := chromeDriver.Start()
	if err != nil {
		log.Println(err)
	}
	desired := webdriver.Capabilities{"Platform": "Mac"}
	required := webdriver.Capabilities{}
	session, err := chromeDriver.NewSession(desired, required)
	if err != nil {
		log.Println(err)
	}

	session.Url("https://www.baidu.com/")
	el, err := session.FindElement("id", "kw")
	if err != nil {
		panic(err)
	}
	fmt.Println(el.Text())

	el.SendKeys("browser")

	el2, err := session.FindElement("id", "su")
	if err != nil {
		panic(err)
	}
	el2.Click()

	time.Sleep(10 * time.Second)
	session.Delete()
	chromeDriver.Stop()
}

//func main() {
	//crawl("http://www.baidu.com/s?wd=%E6%B5%8B%E8%AF%95&rsv_spt=1&
	// rsv_iqid=0xefe8b41d00090434&issp=1&f=3&rsv_bp=0&rsv_idx=2&ie=u
	// tf-8&tn=baiduhome_pg&rsv_enter=0&rsv_sug3=8&rsv_sug1=10&rsv_sug7=100&prefixsug=
	// %E6%B5%8B%E8%AF%95&rsp=9&inputT=35532&rsv_sug4=35534")
	//ago()
	//selen()

	//fed()

//}
