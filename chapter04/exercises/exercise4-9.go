package exercises

import (
	"bufio"
	"os"
)

func wordfreq(f *os.File) map[string]int {
	freq := make(map[string]int)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		str := string(scanner.Bytes())
		freq[str]++
	}
	return freq
}
