package web

import (
	"log"
	"github.com/fedesog/webdriver"
	"fmt"
	"time"
)

func web() {
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

func main() {
	web()
}
