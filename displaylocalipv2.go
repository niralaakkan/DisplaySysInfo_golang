package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	fmt.Println("My first Go program")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			os.Stderr.WriteString("Oops: " + err.Error() + "\n")
			os.Exit(1)
		}

		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				fmt.Fprintf(w, ipnet.IP.String()+"\n")
			}
		}
	})

	log.Fatal(http.ListenAndServe(":8082", nil))
}
