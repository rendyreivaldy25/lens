package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func tablength(thisEntry string) string {
	theseTabs := []string{}
	tabCount := 0
	for i := 0; i < len(thisEntry)/8; i++ {
		tabCount++
	}

	for j := 0; j < tabCount; j++ {
		theseTabs = append(theseTabs, "\t")

	}
	if len(thisEntry) < 8 {
		// println(thisEntry)
		theseTabs = append(theseTabs, "   ")
		theseTabs = append(theseTabs, "\t")
	}

	if len(thisEntry) < 5 {

		theseTabs = append(theseTabs, "\t")
	}

	return strings.Join(theseTabs, "")
}
func thisList(filepath string) {
	f, err := os.Open(filepath)

	er(err)
	files, err := f.Readdir(0)

	er(err)

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})
	// FOLDERS

	for _, entry := range files {
		if entry.IsDir() {
			fmt.Println(boldbrightgreen + entry.Name() + tablength(entry.Name()) + green + "FOLDER" + tablength("FOLDER") + entry.ModTime().String()[:16] + green + " DIR")
		}
		if !entry.IsDir() {
			if entry.Mode().String() == "-rwxr-xr-x" && entry.Mode().Type().String() != "L---------" {
				fmt.Printf("%s%s%s%s%s%s\n", brightyellow, entry.Name(), tablength(entry.Name()), filesize(entry.Size()), tablength(filesize(entry.Size())), entry.ModTime().String()[:16]+yellow+" EXE")
			}

		}
		if entry.Mode().String() != "-rwxr-xr-x" && entry.Mode().Type().String() == "L---------" {
			fmt.Printf("%s%s%s%s%s%s\n", brightmagenta, entry.Name(), tablength(entry.Name()), filesize(entry.Size()), tablength(filesize(entry.Size())), entry.ModTime().String()[:16]+magenta+" SYM")
		}
		if !entry.IsDir() {
			thisSize := filesize(entry.Size())
			if entry.Mode().String() != "-rwxr-xr-x" && entry.Mode().Type().String() != "L---------" {
				fmt.Printf("%s%s%s%s%s%s%s\n", nc, entry.Name(), nc, tablength(entry.Name()), thisSize, tablength(thisSize), entry.ModTime().String()[:16])
			}
		}
	}
}

func filesize(thisFileSize int64) string {
	const byteSize = 1024
	if thisFileSize < byteSize {
		return fmt.Sprintf("%d B", thisFileSize)
	}
	divisor, unitType := int64(byteSize), 0
	for n := thisFileSize / byteSize; n >= byteSize; n /= byteSize {
		divisor *= byteSize
		unitType++
	}
	return fmt.Sprintf("%.2f %cB",
		float64(thisFileSize)/float64(divisor), "KMGTPE"[unitType])
}

func er(err error) {
	if err != nil {
		log.Println(err)
	}
}
