// binary is a tool which parses PE binary executables.
package main

import (
	"flag"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/mewrev/binary/pe"
)

func main() {
	flag.Parse()
	for _, path := range flag.Args() {
		err := parse(path)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// parse parses the provided
func parse(path string) error {
	file, err := pe.Parse(path)
	if err != nil {
		return err
	}
	spew.Dump(file)
	return nil
}
