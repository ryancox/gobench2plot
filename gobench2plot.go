package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {

	// Examples:
	// BenchmarkParsePage        100      17788153 ns/op
	// BenchmarkWrite_4KB_WithIndex       50000         61010 ns/op      67.14 MB/s         598 B/op         16 allocs/op

	nsExp := regexp.MustCompile("^(Benchmark.+?)\\s+.*?(\\S+) ns/op.*")
	allocedBytesExp := regexp.MustCompile("^(Benchmark.+?)\\s+.*?(\\S+) B/op.*")
	allocedExp := regexp.MustCompile("^(Benchmark.+?)\\s+.*?(\\S+) allocs/op.*")
	mbExp := regexp.MustCompile("^(Benchmark.+?)\\s+.*?(\\S+) MB/s.*")

	nsMap := make(map[string]string)
	allocedBytesMap := make(map[string]string)
	allocedMap := make(map[string]string)
	mbMap := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if tokens := nsExp.FindStringSubmatch(line); tokens != nil {
			nsMap[tokens[1]] = tokens[2]
		}
		if tokens := allocedBytesExp.FindStringSubmatch(line); tokens != nil {
			allocedBytesMap[tokens[1]] = tokens[2]
		}
		if tokens := allocedExp.FindStringSubmatch(line); tokens != nil {
			allocedMap[tokens[1]] = tokens[2]
		}
		if tokens := mbExp.FindStringSubmatch(line); tokens != nil {
			mbMap[tokens[1]] = tokens[2]
		}
	}
	fmt.Printf("<Benchmarks>\n")
	fmt.Printf(" <NsPerOp>\n")
	for key := range nsMap {
		fmt.Printf("  <%v>%v</%v>\n", key, nsMap[key], key)
	}
	fmt.Printf(" </NsPerOp>\n")

	fmt.Printf(" <AllocsBytesPerOp>\n")
	for key := range allocedBytesMap {
		fmt.Printf("  <%v>%v</%v>\n", key, allocedBytesMap[key], key)
	}
	fmt.Printf(" </AllocsBytesPerOp>\n")

	fmt.Printf(" <AllocsPerOp>\n")
	for key := range allocedMap {
		fmt.Printf("  <%v>%v</%v>\n", key, allocedMap[key], key)
	}
	fmt.Printf(" </AllocsPerOp>\n")

	fmt.Printf(" <mbPerSec>\n")
	for key := range mbMap {
		fmt.Printf("  <%v>%v</%v>\n", key, mbMap[key], key)
	}
	fmt.Printf(" </mbPerSec>\n")
	fmt.Printf("</Benchmarks>\n")
}
