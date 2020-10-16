// 这个文件是做什么的。
package main

import (
	"fmt"
	"os"
	"bufio"
	"encoding/json"
	"strings"
)

func main() {
	fmt.Println("please input name.")
	inputReader := bufio.NewReader(os.Stdin)
	name, errname := inputReader.ReadString('\n')
	if errname != nil {
		fmt.Println("please input correct name.")
	}
	name = strings.Replace(name, "\n", "", -1)
	fmt.Println("please input address.")
	inputReader = bufio.NewReader(os.Stdin)
	address, erraddress := inputReader.ReadString('\n')
	if erraddress != nil {
		fmt.Println("please input correct address.")
	}
	address = strings.Replace(address, "\n", "", -1)
	p := map[string]string {
		"name": name,
		"address": address,
	}
	b, errb := json.Marshal(p)
	if errb != nil {
		fmt.Println("error", errb)
	}
	fmt.Println(string(b))
}
