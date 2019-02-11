package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var err error

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <filename|fileuri>\n", filepath.Base(os.Args[0]))
	}

	flag.Parse()

	if flag.NArg() > 0 {

		filename := flag.Arg(0)

		var f io.ReadCloser
		var data *http.Response
		if strings.HasPrefix(filename, "http://") || strings.HasPrefix(filename, "https://") {
			data, err = http.Get(filename)
			f = data.Body
		} else {
			f, err = os.Open(filename)
		}
		
		if err != nil {
			panic(fmt.Errorf("Error: %s\n", err))
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
	
		fmt.Println("#EXTM3U")

		for scanner.Scan() {
			l := scanner.Text()
			v := strings.Split(l, ",")
			if (len(v) == 5) {
				fmt.Printf("#EXTINF:-1 tvg-ID=\"%s\" tvg-name=\"%s\" tvg-logo=\"%s\" group-title=\"%s\",%s\n%s\n", v[2], v[1], v[3], v[0], v[1], v[4])
			}
		}

	}
}
