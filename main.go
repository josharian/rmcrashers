package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	log.SetFlags(0)
	if len(os.Args) != 2 {
		log.Fatal("usage: rmcrashers regexp")
	}

	re, err := regexp.Compile(os.Args[1])
	check(err)

	mm, err := filepath.Glob("crashers/*.output")
	check(err)
	for _, m := range mm {
		b, err := ioutil.ReadFile(m)
		check(err)
		x := re.Find(b)
		if x == nil {
			// no match
			continue
		}
		base := m[:len(m)-len(".output")]
		log.Printf("rm crashers/%s.*", base)
		check(os.Remove(m))
		check(os.Remove(base))
		check(os.Remove(base + ".quoted"))
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
