package percentage

import (
	"github.com/ManuchekhrM/convert-number-to-text/models"
	"github.com/ManuchekhrM/convert-number-to-text/valueObjects"
	"github.com/shopspring/decimal"
)

func ConvertNumber(number interface{}, lang string) (string, error) {
	var result string

	switch number.(type) {
	case decimal.Decimal:
		num := number.(decimal.Decimal)
		text, err := ConvNumToTextRUWithPercent(num, lang)
		if err != nil {
			return "", valueObjects.ErrInvalidLanguage
		}
		result = text
	}

	return result, nil
}

func ConvNumToTextRUWithPercent(num decimal.Decimal, lang string) (string, error) {
	var result string

	needLang, ok := models.Language[lang].(map[int]string)
	if !ok {
		return "", valueObjects.ErrInvalidLanguage
	}

	percent := int(num.IntPart())
	decimalPart := num.Sub(decimal.NewFromInt(int64(percent)))

	switch {
	case percent < 10:
		result = needLang[percent]
	case percent > 10 && percent < 20:
		result = needLang[percent-10]
	default:
		tensDigit := percent / 10
		onesDigit := percent % 10
		result = needLang[tensDigit] + " " + needLang[onesDigit]
	}

	if !decimalPart.IsZero() {
		decimalPartString := int(decimalPart.String()[2] - '0')
		decimalText := needLang[decimalPartString]
		result += " целых и " + decimalText + " десятых" + models.Ru
	}

	return result, nil
}

func ConvNumToTextTJWithPercent(per decimal.Decimal) string {
	var result string

	percent := per.IntPart()
	decimalPart := per.Sub(decimal.NewFromInt(percent))

	switch {
	case percent < 10:
		result = models.OneToTenTJ[percent]
	case percent > 10 && percent < 20:
		result = models.TenToNineteenTJ[percent-10]
	default:
		tensDigit := percent / 10
		onesDigit := percent % 10
		result = models.TensTJ[tensDigit] + " " + models.OneToTenTJ[onesDigit]
	}

	if !decimalPart.IsZero() {
		decimalPartString := int(decimalPart.String()[2] - '0')
		decimalText := models.OneToTenTJ[decimalPartString]
		result += " аз даҳ " + decimalText + models.TJ
	}

	return result
}

func ConvNumToTextENWithPercent(per decimal.Decimal) string {
	var result string

	percent := per.IntPart()
	decimalPart := per.Sub(decimal.NewFromInt(percent))

	switch {
	case percent < 10:
		result = models.OneToTenEN[percent]
	case percent > 10 && percent < 20:
		result = models.TenToNineteenEN[percent-10]
	default:
		tensDigit := percent / 10
		onesDigit := percent % 10
		result = models.TensEN[tensDigit] + " " + models.OneToTenEN[onesDigit]
	}

	if !decimalPart.IsZero() {
		decimalPartString := int(decimalPart.String()[2] - '0')
		decimalText := models.OneToTenEN[decimalPartString]
		result += " point " + decimalText + models.EN
	}

	return result
}
