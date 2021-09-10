package main

import (
	_ "embed"
	"fmt"
	"net"
	"os"
	"strings"
)

//go:embed german-cities.txt
var s string

func main() {

	arg := os.Args

	city_list := strings.Fields(s)

	for _, city := range city_list {
		testdomain := arg[1] + "-" + strings.ToLower(city) + ".de"
		nameserver, _ := net.LookupNS(testdomain)
		if len(nameserver) > 0 {
			fmt.Println(testdomain)
		}
	}
}
