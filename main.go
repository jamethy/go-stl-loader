package main

import (
	"fmt"
	"time"
	"github.com/jamethy/go-stl-loader/stlLoaders"
	"os"
	"log"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("program requires at least one argument of filename")
	}

	filePath := os.Args[1]

	start := time.Now()
	stlLoaders.BasicRead(filePath)
	fmt.Println("Execution time: ", time.Since(start))

}
