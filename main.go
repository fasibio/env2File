package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 || os.Args[1] == "--help" {
		fmt.Println("Usage: Command [{Inputfile} {outputfile}]")
		return
	}
	templateFile := os.Args[1]
	buf, err := ioutil.ReadFile(templateFile)
	if err != nil {
		panic(err)
	}
	file := string(buf)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		key := pair[0]
		value := pair[1]
		file = strings.Replace(file, "${{"+key+"}}", value, -1)
	}
	f, err := os.Create(os.Args[2])
	w := bufio.NewWriter(f)
	_, err = w.WriteString(file)
	if err != nil {
		panic(err)
	}
	w.Flush()
	fmt.Println("Write new File")

}
