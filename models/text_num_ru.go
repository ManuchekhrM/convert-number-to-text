package models

var Language = map[string]interface{}{
	"ru": DigitsRu,
	"tj": DigitsTj,
	"en": DigitsEn,
}

type DigitsRU struct {
	OneToTenRU      []string
	TenToNineteenRU []string
	TensRU          []string
	DecimalDigitRU  []string
}

var DigitsRu = DigitsRU{
	OneToTenRU: []string{
		"", "один", "два", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять",
	},
	TenToNineteenRU: []string{
		"десять", "одиннадцать", "двенадцать", "тринадцать", "четырнадцать", "пятнадцать", "шестнадцать", "семнадцать", "восемнадцать", "девятнадцать",
	},
	TensRU: []string{
		"", "десять", "двадцать", "тридцать", "сорок", "пятьдесят", "шестьдесят", "семьдесят", "восемьдесят", "девяносто", "сто",
	},
	DecimalDigitRU: []string{
		"ноль", "одна", "две", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять",
	},
}

type DigitsTJ struct {
	OneToTenTJ      []string
	TenToNineteenTJ []string
	TensTJ          []string
	DecimalDigitTJ  []string
}
var DigitsTj = DigitsTJ{
	OneToTenTJ: []string{
		"", "як", "ду", "се", "чор", "панҷ", "шаш", "ҳафт", "ҳашт", "нӯҳ",
	},
	TenToNineteenTJ: []string{
		"даҳ", "ёздаҳ", "дувоздаҳ", "сенздаҳ", "чордаҳ", "понздаҳ", "шонздаҳ", "ҳабдаҳ", "ҳаждаҳ", "нуздаҳ",
	},
	TensTJ: []string{
		"", "даҳ", "бист", "си", "чил", "панҷоҳ", "шаст", "ҳафтод", "ҳаштод", "навад", "сад",
	},
}

type DigitsEN struct {
	OneToTenEN      []string
	TenToNineteenEN []string
	TensEN          []string
	DecimalDigitEN  []string
}
var DigitsEn = DigitsEN{
	OneToTenEN: []string{
		"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	},
	TenToNineteenEN: []string{
		"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
	},
	TensEN: []string{
		"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety", "one hundred",
	},
}
