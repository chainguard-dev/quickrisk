package main

import (
	"flag"
	"log"
	"os"

	"github.com/chainguard-dev/quickrisk/pkg/format"
	"github.com/chainguard-dev/quickrisk/pkg/quickrisk"
)

func main() {
	csvOutput := flag.Bool("csv", false, "Output results in CSV format")
	dotOutput := flag.Bool("dot", false, "Output results in Graphviz DOT format")
	otmOutput := flag.Bool("otm", false, "Output results in Open Threat Modeling (OTM) format - EXPERIMENTAL")
	threagileOutput := flag.Bool("threagile", false, "Output results in Threagile format - EXPERIMENTAL")
	// defaultRisk := flag.Int("default-impact", 3, "Default impact for risks")
	// defaultLikelihood := flag.Int("default-likelihood", 3, "Default likelihood for risks")
	flag.Parse()

	if flag.NArg() < 1 {
		log.Fatalf("Usage: %s [--csv] [--dot] [--otm] [--threagile] <yaml_file_or_directory>...", os.Args[0])
	}

	c, err := quickrisk.LoadConfigs(flag.Args())
	if err != nil {
		log.Fatalf("Load failed: %v", err)
	}

	if *csvOutput {
		format.CSV(c)
	} else if *dotOutput {
		format.DOT(c)
	} else if *otmOutput {
		format.OTM(c)
	} else if *threagileOutput {
		format.Threagile(c)
	} else {
		format.Text(c)
	}
}
