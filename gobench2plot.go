package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const (
	// BenchmarkParsePage        100      17788153 ns/op
	benchPat = "^(Benchmark.+?)\\s+(\\d+?)\\s+(\\d+?) ns/op"
)

func main() {

	bexp := regexp.MustCompile(benchPat)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("# gobench2plot output - http://github.com/ryancox/gobench2plot\n")
	fmt.Printf("# <Benchmark>=<ns/op>\n")

	for scanner.Scan() {
		line := scanner.Text()
		if tokens := bexp.FindStringSubmatch(line); tokens != nil {
			fmt.Printf("%v=%v\n", tokens[1], tokens[2])
		}
	}
}
