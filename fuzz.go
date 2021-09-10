package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	arg := os.Args

	current_dir, err1 := os.Getwd()
	if err1 != nil {
		fmt.Println("error getting current dir")
	}

	city_list, err2 := os.Open(current_dir + "/german-cities.txt")
	if err2 != nil {
		fmt.Println("error loading city-list")
	}

	scanner := bufio.NewScanner(city_list)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		testdomain := arg[1] + "-" + strings.ToLower(scanner.Text()) + ".de"
		nameserver, _ := net.LookupNS(testdomain)
		if len(nameserver) > 0 {
			fmt.Println(testdomain)
		}
	}

	city_list.Close()
}
