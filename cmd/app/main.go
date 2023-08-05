package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ptakpatryk/linkparser/cmd/parser"
)

func checkError(err error) {
  if(err != nil) {
    fmt.Println(err)
    os.Exit(1)
  }
}

func main() {
	htmlFilePath := flag.String("html", "assets/ex1.html", "path to the html file you want to parse")
  flag.Parse()
  f, err := os.Open(*htmlFilePath)
  checkError(err)

  defer f.Close()

	parser.ParseLink()

}
