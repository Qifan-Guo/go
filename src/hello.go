package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports chan int, results chan int, IP_address string) {
	for p := range ports {
		address := fmt.Sprintf(IP_address+":%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	fmt.Println("Hello , Black Hat Gophers!")
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results, "127.0.0.1")
	}
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(results)
	close(ports)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("Port : %d is open\n", port)
	}

}
