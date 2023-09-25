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
		text, err := ConvNumToTextRuDecimal(num, lang)
		if err != nil {
			return "", valueObjects.ErrInvalidLanguage
		}
		result = text
	case int:
		num := number.(int)
		text, err := ConvNumToTextRuInt(num, lang)
		if err != nil {
			return "", valueObjects.ErrInvalidLanguage
		}
		result = text
	}

	return result, nil
}

func ConvNumToTextRuDecimal(num decimal.Decimal, lang string) (string, error) {
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

func ConvNumToTextRuInt(num int, lang string) (string, error) {
	var result string

	needLang, ok := models.Language[lang].(map[int]string)
	if !ok {
		return "", valueObjects.ErrInvalidLanguage
	}

	percent := num

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

	return result, nil
}
