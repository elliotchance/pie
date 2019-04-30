package pie

type currency struct {
	NumericCode, Exponent int
}

//go:generate pie currencies.*
type currencies map[string]currency

var isoCurrencies = currencies{
	"AUD": {36, -2},
	"USD": {840, -2},
}
