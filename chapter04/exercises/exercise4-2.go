package exercises

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	var flag384, flag512 bool
	flag.BoolVar(&flag384, "sha384", false, "use SHA384 function")
	flag.BoolVar(&flag512, "sha512", false, "use SHA512 function")
	flag.Parse()

	input := bufio.NewReader(os.Stdin)
	fmt.Println("Type something: ")
	data, err := input.ReadBytes('\n')
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	if flag384 {
		fmt.Printf("SHA384 checksum: %x\n", sha512.Sum384(data))
	} else if flag512 {
		fmt.Printf("SHA512 checksum: %x\n", sha512.Sum512(data))
	} else {
		fmt.Printf("SHA256 checksum: %x\n", sha256.Sum256(data))
	}
}
