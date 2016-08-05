package browser

import (
	"github.com/fedesog/webdriver"
	"fmt"
)

type BrowserSession struct {
	*webdriver.Session
}

func (session BrowserSession) Fill_by_id(id string, value string) {
	el, err := session.FindElement("id", id)
	if err != nil {
		panic(err)
	}
	fmt.Println(el.Text())

	el.SendKeys(value)
}
