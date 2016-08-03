package main

import (
	"fmt"
	"sourcegraph.com/sourcegraph/go-selenium"
)

func main() {

	var webDriver selenium.WebDriver
	var err error
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	if webDriver, err = selenium.NewRemote(caps, "http://testerp.dmall.com/login"); err != nil {
		fmt.Printf("Failed to open session: %s\n", err)
		return
	}
	defer webDriver.Quit()

	//err = webDriver.Get("https://www.baidu.com")
	//if err != nil {
	//	fmt.Printf("Failed to load page: %s\n", err)
	//	return
	//}

	//if title, err := webDriver.Title(); err == nil {
	//	fmt.Printf("Page title: %s\n", title)
	//} else {
	//	fmt.Printf("Failed to get page title: %s", err)
	//	return
	//}

	//var elem selenium.WebElement
	//elem, err = webDriver.FindElement(selenium.ByCSSSelector, ".repo .name")
	//if err != nil {
	//	fmt.Printf("Failed to find element: %s\n", err)
	//	return
	//}
	//
	//if text, err := elem.Text(); err == nil {
	//	fmt.Printf("Repository: %s\n", text)
	//} else {
	//	fmt.Printf("Failed to get text of element: %s\n", err)
	//	return
	//}
}
