package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"io/ioutil"
)


func main() {

	var buf []byte
	var err error

	flag.Parse()

	switch narg := flag.NArg(); {

		case  narg == 0:
			buf, err = ioutil.ReadAll(os.Stdin)
			if err != nil {
				panic(fmt.Errorf("Error: %s\n", err))
			}

		case narg > 0:
			buf, err = ioutil.ReadFile(flag.Arg(0))
			if err != nil {
				panic(fmt.Errorf("Error: %s\n", err))
			}

	}

	uppercase := strings.ToUpper(string(buf))
	fmt.Printf("%v\n", uppercase)
}
