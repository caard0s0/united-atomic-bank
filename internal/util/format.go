package util

import (
	"time"

	"github.com/bojanz/currency"
)

func FormatDate(date time.Time) string {
	now := date
	formattedDate := now.Format("Jan 02, 2006 15:04")

	return formattedDate
}

func FormatCurrency(amountTransfer string, currencyCode string) string {
	locale := currency.NewLocale("US")
	formatter := currency.NewFormatter(locale)
	amount, _ := currency.NewAmount(amountTransfer, currencyCode)
	formattedCurrency := formatter.Format(amount)

	return formattedCurrency
}
