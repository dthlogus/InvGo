package models

type StockData struct {
	PriceToBook                 float32 `json:"priceToBook"`
	RegularMarketPrice          float32 `json:"regularMarketPrice"`
	TrailingAnnualDividendYield float32 `json:"trailingAnnualDividendYield"`
	EpsTrailingTwelveMonths     float32 `json:"epsTrailingTwelveMonths"`
	BookValue                   float32 `json:"bookValue"`
	Symbol                      string  `json:"symbol"`
	LongName                    string  `json:"longName"`
}

type StockDTO struct {
	PVP      float32 `json:"pvp"`
	PV       float32 `json:"pv"`
	DY       float32 `json:"dy"`
	VPA      float32 `json:"vpa"`
	LPA      float32 `json:"lpa"`
	Symbol   string  `json:"symbol"`
	LongName string  `json:"longName"`
}

func ToDTOs(data []StockData) []StockDTO {
	dtos := make([]StockDTO, len(data))
	for i, d := range data {
		dtos[i] = StockDTO{
			PVP:      d.RegularMarketPrice / d.BookValue,
			PV:       1 / d.PriceToBook,
			DY:       d.TrailingAnnualDividendYield,
			VPA:      d.BookValue,
			LPA:      d.EpsTrailingTwelveMonths,
			Symbol:   d.Symbol,
			LongName: d.LongName,
		}
	}
	return dtos
}
