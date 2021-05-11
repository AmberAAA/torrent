package main

import (
	"flag"
	"fmt"
	"os"
)

const help = `
-h		help
-e		env
`

func main() {
	fmt.Println(help)
	for _, v := range os.Args {
		fmt.Println(v)
	}
	var port int
	flag.IntVar(&port, "p", 8000, "Specify port to use. default is 8000.")
	flag.Parse()

	fmt.Printf("port = %d", port)
}
