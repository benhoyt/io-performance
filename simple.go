// Run with:
// go build simple.go
// sudo sysctl vm.drop_caches=3; ./simple <kjvbible_x100.txt >/dev/null

package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	tStart := time.Now()
	content, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	tRead := time.Now()

	words := strings.Fields(strings.ToLower(string(content)))
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}
	tProcess := time.Now()

	var ordered []Count
	for word, count := range counts {
		ordered = append(ordered, Count{word, count})
	}
	sort.Slice(ordered, func(i, j int) bool {
		return ordered[i].Count > ordered[j].Count
	})
	tSort := time.Now()

	for _, count := range ordered {
		fmt.Println(string(count.Word), count.Count)
	}
	tEnd := time.Now()

	fmt.Fprintf(os.Stderr, "Reading   : %v\n", tRead.Sub(tStart).Seconds())
	fmt.Fprintf(os.Stderr, "Processing: %v\n", tProcess.Sub(tRead).Seconds())
	fmt.Fprintf(os.Stderr, "Sorting   : %v\n", tSort.Sub(tProcess).Seconds())
	fmt.Fprintf(os.Stderr, "Outputting: %v\n", tEnd.Sub(tSort).Seconds())
	fmt.Fprintf(os.Stderr, "TOTAL     : %v\n", tEnd.Sub(tStart).Seconds())
}

type Count struct {
	Word  string
	Count int
}
