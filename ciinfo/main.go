package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/gkampitakis/ciinfo"
)

const (
	YES = 0
	NO  = -1
)

func main() {
	isPr := flag.Bool("pr", false, "check if shell is running on CI for a Pull Request.")
	output := flag.String("output", "", "you can output info [json, pretty].")

	flag.Parse()

	run(os.Exit, os.Stdout, *isPr, *output)
}

func run(exit func(int), w io.Writer, isPR bool, output string) {
	result := NO

	if isPR {
		if ciinfo.IsPr {
			result = YES
		}

		exit(result)
	}

	if output == "json" {
		jsonPrint(w)
		return
	}

	if output == "pretty" {
		prettyPrint(w)
		return
	}

	if ciinfo.IsCI {
		result = YES
	}

	exit(result)
}

func jsonPrint(w io.Writer) {
	fmt.Fprintf(w, `{
	"is_ci": %t,
	"ci_name": "%s",
	"pull_request": %t
}
`, ciinfo.IsCI, ciinfo.Name, ciinfo.IsPr)
}

func prettyPrint(w io.Writer) {
	if !ciinfo.IsCI {
		fmt.Fprintln(w, "Not running on CI.")
		return
	}

	if ciinfo.Name != "" {
		fmt.Fprintf(w, "CI Name: %s\n", ciinfo.Name)
	} else {
		fmt.Fprintln(w, "Running on CI.")
	}

	if ciinfo.IsPr {
		fmt.Fprintln(w, "Is Pull Request.")
	}
}
