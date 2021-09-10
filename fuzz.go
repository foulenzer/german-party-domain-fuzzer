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

	current_dir, err_dir := os.Getwd()
	if err_dir != nil {
		fmt.Println("error getting current dir")
	}

	city_list, err_list := os.Open(current_dir + "/german-cities.txt")
	if err_list != nil {
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
