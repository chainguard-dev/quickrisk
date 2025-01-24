package main

import (
	"flag"
	"log"
	"os"

	"github.com/chainguard-dev/quickrisk/pkg/format"
	"github.com/chainguard-dev/quickrisk/pkg/quickrisk"
)

func main() {
	formatFlag := flag.String("format", "text", "Output format (csv, dot, otm [EXPERIMENTAL], threagile [EXPERIMENTAL])")
	outputFlag := flag.String("output", "", "Output file (default: stdout)")
	flag.Parse()

	if flag.NArg() < 1 {
		log.Fatalf("Usage: %s [--format=whatever] [--output=file] <yaml_file_or_directory>...", os.Args[0])
	}

	// Open the output file or use stdout
	var of *os.File
	var err error
	if *outputFlag != "" {
		of, err = os.Create(*outputFlag)
		if err != nil {
			log.Fatalf("Failed to create output file: %v", err)
		}
		defer of.Close()
	} else {
		of = os.Stdout
	}

	// Load configurations
	c, err := quickrisk.LoadConfigs(flag.Args())
	if err != nil {
		log.Fatalf("Load failed: %v", err)
	}

	errs := quickrisk.Validate(c)
	if len(errs) > 0 {
		log.Printf("validation failed:")
		for _, e := range errs {
			log.Printf("* %s", e)
		}
	}

	// Handle the output format
	switch *formatFlag {
	case "csv":
		format.CSV(of, c)
	case "dot":
		format.DOT(of, c)
	case "png":
		format.PNG(of, c)
	case "otm":
		format.OTM(of, c)
	case "threagile":
		format.Threagile(of, c)
	default:
		format.Text(of, c)
	}
}
