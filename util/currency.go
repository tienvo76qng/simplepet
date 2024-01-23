package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

// is support currency return true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}
