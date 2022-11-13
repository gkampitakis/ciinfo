package main

import (
	"fmt"
	"os"
)

const template = `
// CI is running on %s
var %s = vendorsIsCI["%s"]
`

func main() {
	f, err := os.Create("constants.go")
	if err != nil {
		fmt.Printf("error opening file %s\n", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(
		"package ciinfo\n\n// This file is generated. Run `make compile-constants` to update.\n",
	); err != nil {
		fmt.Printf("error writing on file %s\n", err)
		return
	}

	for _, v := range vendors {
		if _, err := fmt.Fprintf(f, template, v.name, v.constant, v.constant); err != nil {
			fmt.Printf("error writing on file %s\n", err)
			return
		}
	}
}
