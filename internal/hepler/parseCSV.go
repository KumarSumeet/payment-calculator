package hepler

import (
	bs "cursol/internal/assets/bankstatements"
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

// ParseCSV parses from csv
func ParseCSV(filepath string) (bankstatements []bs.BankStatement, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var statements []bs.BankStatement
	reader := csv.NewReader(file)
	reader.Read()

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		credit, _ := strconv.ParseFloat(line[7], 64)
		debit, _ := strconv.ParseFloat(line[8], 64)

		b := bs.BankStatement{
			Date:       line[0],
			Narrative1: line[1],
			Narrative2: line[2],
			Narrative3: line[3],
			Narrative4: line[4],
			Narrative5: line[5],
			Type:       line[6],
			Credit:     credit,
			Debit:      debit,
			Currency:   line[9],
		}

		statements = append(statements, b)
	}

	return statements, nil
}
