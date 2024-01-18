package assets

// BankStatement is the BankStatement struct
type BankStatement struct {
	Date       string
	Narrative1 string
	Narrative2 string
	Narrative3 string
	Narrative4 string
	Narrative5 string
	Type       string
	Credit     float64
	Debit      float64
	Currency   string
}
