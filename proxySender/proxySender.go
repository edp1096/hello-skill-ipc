package main // import "proxySender"

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	wordPtr := flag.String("cmd", "notmycmmand", "a string")
	flag.Parse()

	fmt.Println("cmd", *wordPtr)

	_, err := http.Get("http://localhost:1323")
	if err != nil {
		log.Fatalln(err)
	}
}
