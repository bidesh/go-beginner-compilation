package main

import (
	"fmt"
	"os"
)

func main() {
	{
		if len(os.Args) > 1 {
			fmt.Println("I am the guest")
		} else {
			fmt.Println("I am the host")
		}
	}
}

//goimports to import the repo's required
//gofmt to tidy the code
//testing123