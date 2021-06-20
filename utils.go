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
func thisList(filepath string, args string) {
	f, err := os.Open(filepath)

	er(err)
	files, err := f.Readdir(0)

	er(err)

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})
	// FOLDERS

	indexToHide := strings.Split(args, ",")

	for _, entry := range files {

		if entry.IsDir() {
			fmt.Println(getColor(checkIsColumnHide(indexToHide, "color"), "DIR") + entry.Name() + tablength(entry.Name()) + getColor(checkIsColumnHide(indexToHide, "color"), "TYPEDIR") + "FOLDER" + tablength("FOLDER") + entry.ModTime().String()[:16] + getColor(checkIsColumnHide(indexToHide, "color"), "TYPEDIR") + " DIR")
		}
		if !entry.IsDir() {
			if entry.Mode().String() == "-rwxr-xr-x" && entry.Mode().Type().String() != "L---------" {
				fmt.Printf("%s%s%s%s%s%s\n", getColor(checkIsColumnHide(indexToHide, "color"), "EXE"), entry.Name(), tablength(entry.Name()), filesize(entry.Size()), tablength(filesize(entry.Size())), entry.ModTime().String()[:16]+getColor(checkIsColumnHide(indexToHide, "color"), "TYPEEXE")+" EXE")
			}

		}
		if entry.Mode().String() != "-rwxr-xr-x" && entry.Mode().Type().String() == "L---------" {
			fmt.Printf("%s%s%s%s%s%s\n", getColor(checkIsColumnHide(indexToHide, "color"), "SYM"), entry.Name(), tablength(entry.Name()), filesize(entry.Size()), tablength(filesize(entry.Size())), entry.ModTime().String()[:16]+getColor(checkIsColumnHide(indexToHide, "color"), "TYPESYM")+" SYM")
		}
		if !entry.IsDir() {
			thisSize := filesize(entry.Size())
			if entry.Mode().String() != "-rwxr-xr-x" && entry.Mode().Type().String() != "L---------" {
				fmt.Printf("%s%s%s%s%s%s%s\n", nc, entry.Name(), nc, tablength(entry.Name()), thisSize, tablength(thisSize), entry.ModTime().String()[:16])
			}
		}
	}
}

func checkIsColumnHide(args []string, column string) bool {
	for _, data := range args {
		if strings.Trim(data, " ") == column {
			return true
		}
	}
	return false
}

func getColor(colorArgumentFlag bool, dataType string) string {
	if !colorArgumentFlag {
		switch dataType {
		case "DIR":
			return boldbrightgreen
		case "EXE":
			return brightyellow
		case "SYM":
			return brightmagenta
		case "TYPEDIR":
			return green
		case "TYPEEXE":
			return yellow
		case "TYPESYM":
			return magenta
		default:
			return nc
		}
	}
	return nc
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

	return fmt.Sprintf("%.1f %cB",
		float64(thisFileSize)/float64(divisor), "KMGTPE"[unitType])
}

func er(err error) {
	if err != nil {
		log.Println(err)
	}
}
