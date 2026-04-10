package main

import (
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

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
	for name := range sheetMap {
		fmt.Println(name)
	}
}
