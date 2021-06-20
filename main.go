package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	hideArgs := flag.String("hide", "- empty -", "define which things to hide, i.e : color, size, date, type. separated with comma")
	flag.Parse()
	if len(flag.Args()) == 0 {
		thisList(filepath, *hideArgs)
	} else {
		for i, j := range flag.Args() {
			if i != 0 {
				fmt.Println(brightyellow + "Directory:\t" + j + nc)
				thisList(j, *hideArgs)
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
