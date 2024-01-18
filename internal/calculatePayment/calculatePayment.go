package calculatePayment

import (
	bs "cursol/internal/assets/bankstatements"
	"strings"
)

// CalcPayments calculates the respective payments
func CalcPayments(statements []bs.BankStatement, date string) map[string]float64 {
	totals := make(map[string]float64)
	for _, st := range statements {
		if st.Date != date {
			continue
		}

		if strings.Contains(st.Narrative1, "PAY") || strings.Contains(st.Narrative2, "PAY") || strings.Contains(st.Narrative3, "PAY") || strings.Contains(st.Narrative4, "PAY") || strings.Contains(st.Narrative5, "PAY") {
			totals[st.Currency] += st.Debit
		}
	}
	return totals
}
