package main

import (
	"github.com/fedesog/webdriver"
	"log"
	"fmt"
	"black-widow/src/aaa"
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

func main() {
	session, err := Open_url("http://www.baiduyun.me/forum.php");
	fmt.Println(err)
	fmt.Println(session)

	browser_sesion := browser.BrowserSession{session}

	browser_sesion.Fill_by_id("ls_username", "hehe198504")
}
