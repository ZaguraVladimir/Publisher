package main

import (
	"Publisher/publisher"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var sourceDir = flag.String("s", "Source books", "Каталог и исходными текстами книг в формате MD")
var resultDir = flag.String("r", "Result books", "Каталог, в который будут помещен результат")

func init() {
	flag.Parse()

	workDir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	*sourceDir = filepath.Join(workDir, *sourceDir)
	*resultDir = filepath.Join(workDir, *resultDir)
}

func main() {

	p := publisher.NewPublisher(*sourceDir, *resultDir)
	fmt.Print(p)

}
