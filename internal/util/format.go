package util

import (
	"strconv"
	"time"

	"github.com/bojanz/currency"
)

func FormatDate(date time.Time) string {
	now := date
	formattedDate := now.Format("Jan 02, 2006 15:04")

	return formattedDate
}

func FormatCurrency(amountTransfer int64, currencyCode string) string {
	amountTransferToString := strconv.Itoa(int(amountTransfer))

	locale := currency.NewLocale("US")
	formatter := currency.NewFormatter(locale)
	amount, _ := currency.NewAmount(amountTransferToString, currencyCode)
	formattedCurrency := formatter.Format(amount)

	return formattedCurrency
}
