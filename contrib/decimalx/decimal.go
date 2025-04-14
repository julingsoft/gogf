package decimalx

import (
	"github.com/shopspring/decimal"
)

// IntToDecimal decimal转浮点数
func IntToDecimal(f int64) decimal.Decimal {
	return decimal.NewFromInt(f)
}

// FloatToDecimal 浮点数转decimal
func FloatToDecimal(f float64) decimal.Decimal {
	return decimal.NewFromFloat(f)
}

// FloatToInt 浮点数转整数
// f: 浮点数，比如元
// p: 精度，比如分，精度为2，那么1.00000001会被转为100
func FloatToInt(f float64, p int64) int64 {
	return FloatToDecimal(f).Mul(IntToDecimal(p)).IntPart()
}

// YuanToFen 元转分
// yuan: 以元为单位的浮点数
func YuanToFen(yuan float64) int64 {
	return FloatToInt(yuan, 100)
}

// FenToYuan 分转元
// fen: 以分为单位的整数
func FenToYuan(fen int64) float64 {
	decimalFen := IntToDecimal(fen)
	decimalYuan := decimalFen.Div(IntToDecimal(100))
	return decimalYuan.InexactFloat64()
}
