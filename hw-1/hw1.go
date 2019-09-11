package main

import (
	"fmt"
	"log"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatal("error: ", err)
	}
	fmt.Println(time)
	os.Exit(0)
}
