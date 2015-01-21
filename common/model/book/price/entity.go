package price

import "github.com/opbk/openbook/common/model/book/price/pricetype"

type Price struct {
	pricetype.Type
	Price float64
}
