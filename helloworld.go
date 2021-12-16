package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	//fmt.Println("Hello World")
	WhatToSay := doctor.Intro()
	fmt.Println(WhatToSay)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("->")
		UserInput, _ := reader.ReadString('\n')
		UserInput = strings.Replace(UserInput, "\r\n", "", -1)
		UserInput = strings.Replace(UserInput, "\n", "", -1)
		if UserInput == "quit" {
			break
		} else {
			Response := doctor.Response(UserInput)
			fmt.Println(Response)

		}
	}

}
