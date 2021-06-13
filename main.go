package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		thisList(filepath)
	} else if len(os.Args) == 2 {
		thisList(os.Args[1])
	} else if len(os.Args) > 2 {
		for i, j := range os.Args {
			if i != 0 {

				fmt.Println(brightyellow + "Directory:\t" + j + nc)
				thisList(j)
				if i < len(os.Args) {
					fmt.Println()
				}
			}
		}
	}

}

func init() {
	osCheck()

}
