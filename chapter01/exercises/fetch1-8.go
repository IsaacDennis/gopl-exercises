package main
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)
func hasHttp(str string) bool {
	return strings.HasPrefix(str, "http://")
}
func main(){
	for _, url := range os.Args[1:] {
		if !hasHttp(url) { url = "http://" + url }
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//Padrão
		b, err := ioutil.ReadAll(resp.Body);
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
