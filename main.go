package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ProFL/gophercises-link/link"
	"golang.org/x/net/html"
)

func main() {
	filePath := flag.String("file", "", "path to the html file")
	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatalf("Failed to open file at %s: %v\n", *filePath, err.Error())
	}
	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			log.Panicln("Failed to close html file", closeErr.Error())
		}
	}()

	rootNode, err := html.Parse(file)
	if err != nil {
		log.Fatalln("Failed to parse HTML", err.Error())
	}

	links := link.ParseLinksFromRoot(rootNode)
	fmt.Printf("%+v\n", links)
}
