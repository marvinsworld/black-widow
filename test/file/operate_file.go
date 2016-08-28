package main

import (
	"fmt"
	_ "black-widow/test/web"
)

func init(){
	fmt.Println("init")
}

func main() {
	//web.Test1()
	fmt.Println("Hello, 世界")
}
