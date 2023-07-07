package exercises

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[*os.File]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, nestedMap := range counts {
		for file, n := range nestedMap {
			if n > 1 {
				fmt.Printf("'%s' line has %d duplicates in file %s\n", line, n, file.Name())
			}
		}
	}
}
func countLines(f *os.File, counts map[string]map[*os.File]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if counts[line] == nil {
			counts[line] = make(map[*os.File]int)
		}
		counts[line][f]++
	}

}
