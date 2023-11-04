package main

import (
	"fmt"
	"log"
	"os"

	"github.com/deut/garage-accounting/db"
)

func main() {

	err := db.Connect()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	fmt.Println(db.InitializeSchema())
}
