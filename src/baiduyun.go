package main

import (
	"github.com/fedesog/webdriver"
	"log"
	"fmt"
	"black-widow/src/browser"
	"time"
)

func Open_url(url string) (*webdriver.Session, error) {
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
	session.Url(url)
	return session, err
}

func Fill_by_id(session *webdriver.Session, id string, value string) {
	el, err := session.FindElement("id", id)
	if err != nil {
		panic(err)
	}
	fmt.Println(el.Text())

	el.SendKeys(value)
}

func Open_newTab(url string) {

}

func main() {
	session, err := Open_url("http://www.baiduyun.me/forum.php");
	fmt.Println(err)
	fmt.Println(session)

	browser_sesion := browser.BrowserSession{session}

	browser_sesion.Fill_by_id("ls_username", "hehe198504")
	browser_sesion.Fill_by_id("ls_password", "hehe1234567")

	el2, err := session.FindElement("class name", "pn")
	if err != nil {
		panic(err)
	}
	el2.Click()
	time.Sleep(2 * time.Second)
	//http://www.baiduyun.me/member.php?mod=logging&action=login&loginsubmit=yes&infloat=yes&lssubmit=yes&inajax=1

	//el3, err := session.FindElement("id", "header-loggin-btn")
	//if err != nil {
	//	panic(err)
	//}
	//el3.Submit()

	//el4, err := session.FindElement("class name", "gt_slider_knob")
	//if err != nil {
	//	panic(err)
	//}
	//el4.Submit()


	//fmt.Println("开始打卡")
	//time.Sleep(6 * time.Second)
	//fmt.Println("开始")
	//daka, err := session.FindElement("id", "pper_a")
	//if err != nil {
	//	panic(err)
	//}
	//daka.Click()


	el, err := session.FindElement("tag name", "body")
	if err != nil {
		panic(err)
	}
	fmt.Println(el.Text())

	err = el.SendKeys("ctrl t")

}
