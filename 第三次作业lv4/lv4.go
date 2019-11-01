package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	newfile  *os.File
	err  error
	filieName = "proverb.txt"
)

func main() {
	newfile, err = os.Create(filieName)
	if err != nil {
		log.Fatal(err)
	}
	defer newfile.Close()
	newfile.Write([]byte("don't communicate by sharing memory share memory by communicating"))
	buf,_ :=ioutil.ReadFile("proverb.txt")
	fmt.Println(string(buf))
}
