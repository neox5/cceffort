package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"

	"github.com/xuri/excelize/v2"
)

type Sheet struct {
	ID   int
	Name string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input.xlsx>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	f, err := excelize.OpenFile(inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	sheetMap := f.GetSheetMap()
	var cwSheets []Sheet

	re := regexp.MustCompile(`^KW\d{2}$`)

	for id, name := range sheetMap {
		if re.MatchString(name) {
			cwSheets = append(cwSheets, Sheet{ID: id, Name: name})
		}
	}

	// Sort by sheet name (alphabetically)
	sort.Slice(cwSheets, func(i, j int) bool {
		return cwSheets[i].Name < cwSheets[j].Name
	})

	for _, sheet := range cwSheets {
		fmt.Println(sheet.Name)
	}
}
