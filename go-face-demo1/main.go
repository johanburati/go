//
// Detect and name faces on a group image, then check if we can find the target face in that group
//

package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/fatih/color"
	"github.com/Kagami/go-face"
	flags "github.com/jessevdk/go-flags"
)

type Options struct {
	TargetImage string `short:"t" long:"target" description:"Target image (single face)" required:"true"`
	GroupImage string `short:"g" long:"group" description:"Group image (multiple faces)" required:"true"`
	Names string `short:"n" long:"names" description:"Names (delimited by comma)" required:"true"`
}

func main() {

	var opts Options

	parser := flags.NewParser(&opts, flags.HelpFlag)
	args, err := parser.Parse()
	if err != nil {
		log.Fatalf("Error: %s (%s)\n", err, args)
	}
	color.Set(color.FgGreen)
	fmt.Printf("# GO-FACE DEMO #1\n")

	// Init the recognizer. with .dat files
	rec, err := face.NewRecognizer(".")
	if err != nil {
		log.Fatalf("Can't init face recognizer: %v", err)
	}
	defer rec.Close()

	color.Set(color.FgYellow)
	fmt.Printf("## Detect faces on %s\n", opts.GroupImage)
	names := strings.Split(opts.Names, ",")

	// Recognize faces on that image.
	faces, err := rec.RecognizeFile(opts.GroupImage)
	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}

	if len(faces) != len(names) {
		log.Fatalf("Wrong number of faces")
	}

	// images for each person to get better classification results
	// but in our example we just get them from one big image.
	var samples []face.Descriptor
	var cats []int32
	for i, f := range faces {
		fmt.Printf("- Detected: %s\n", names[i])
		samples = append(samples, f.Descriptor)
		// Each face is unique on that image so goes to its own
		// category.
		cats = append(cats, int32(i))
	}
	// Pass samples to the recognizer.
	rec.SetSamples(samples, cats)

	color.Set(color.FgCyan)
	fmt.Printf("## Recognize face on %s\n", opts.TargetImage)
	targetFace, err := rec.RecognizeSingleFile(opts.TargetImage)
	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}

	if targetFace == nil {
		log.Fatalf("Not a single face on the image")
	}
	match := rec.Classify(targetFace.Descriptor)
	if match < 0 {
		log.Fatalf("Can't classify")
	}
	// Finally print the classified label. 
	fmt.Printf("- Found: %s\n",names[match])
	color.Unset()
}
