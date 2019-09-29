package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jamesnetherton/m3u"
)

// getValues returns the tag values sorted by tag names
func getValues(itags []m3u.Tag) []string {
			names := []string{"group-title", "tvg-name", "tvg-ID", "tvg-logo"}
			
			// create a map out of track.Tags 
			tags := make(map[string]string)
			for _, t := range itags {
				// replace comma with colon, since comma is the fild delimiter in csv files
				v := strings.Replace(t.Value, ",", ":", -1)
				tags[t.Name] = v
			}

			// extract the values in order
			values := make([]string, 0, len(names))
			for _, n := range names {
				values = append(values, tags[n])
			}
			return values
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <filename|fileuri>\n", filepath.Base(os.Args[0]))
	}

	flag.Parse()

	if flag.NArg() > 0 {

		playlist, err := m3u.Parse(flag.Arg(0))

		if err == nil {
			for _, track := range playlist.Tracks {
				values := getValues(track.Tags)
				if len(values) == 0 {
					continue
				}

				// if tvg-name is empty, use the track Name (at the end of the line after the comma)
				if values[1] == "" { values[1] = track.Name }
				values = append(values, track.URI)
				// track.Name is mess up if there is comma in the title, use tvg-name tag instead

				fmt.Println(strings.Join(values, ","))
			}	
		} else {
			fmt.Println(err)
		}
	}
}
