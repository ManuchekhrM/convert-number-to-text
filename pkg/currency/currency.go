package currency

// import (
// 	"github.com/ManuchekhrM/convert-number-to-text/models"
// 	"github.com/shopspring/decimal"
// )

// func ConvNumToTextRuWithPercent(per decimal.Decimal) string {
// 	var result string

// 	percent := per.IntPart()
// 	decimalPart := per.Sub(decimal.NewFromInt(percent))

// 	switch {
// 	case percent < 10:
// 		result = models.OneToTenRU[percent]
// 	case percent > 10 && percent < 20:
// 		result = models.TenToNineteenRU[percent-10]
// 	default:
// 		tensDigit := percent / 10
// 		onesDigit := percent % 10
// 		result = models.TensRU[tensDigit] + " " + models.OneToTenRU[onesDigit]
// 	}

// 	if !decimalPart.IsZero() {
// 		decimalPartString := int(decimalPart.String()[2] - '0')
// 		decimalText := models.DecimalDigitRU[decimalPartString]
// 		result += " целых и " + decimalText + " десятых"
// 	}

// 	return result
// }
