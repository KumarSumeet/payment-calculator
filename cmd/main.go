package main

import (
	calculate "cursol/internal/calculatePayment"
	payment "cursol/internal/hepler"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go csv-file path")
		os.Exit(1)
	}

	filePath := os.Args[1]

	absolutePath, err := filepath.Abs(filePath)
	fmt.Println("Absolute Path:", absolutePath)

	statements, err := payment.ParseCSV(absolutePath)

	if err != nil {
		log.Fatal(err)
	}

	totals := calculate.CalcPayments(statements, "06/03/2011")

	for currency, total := range totals {
		fmt.Printf("%s %.2f\n", currency, total)
	}
}
