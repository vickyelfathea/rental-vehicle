package main

import (
	"log"
	"os"
	"carRent/src/configs/command"
	)
	
func main() {
	if err := command.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

// func sampleHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello from user"))
// }
