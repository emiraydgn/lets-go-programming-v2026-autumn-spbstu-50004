package cbr

import (
	"encoding/xml"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type valCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []valute `xml:"Valute"`
}

type valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

func Load(path string) ([]Currency, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open XML file: %w", err)
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var data valCurs
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("decode XML: %w", err)
	}

	currencies := make([]Currency, 0, len(data.Valutes))

	for _, item := range data.Valutes {
		numCode := 0

		if item.NumCode != "" {
			numCode, err = strconv.Atoi(item.NumCode)
			if err != nil {
				return nil, fmt.Errorf("invalid NumCode %q: %w", item.NumCode, err)
			}
		}

		valueText := strings.ReplaceAll(item.Value, ",", ".")
		value, err := strconv.ParseFloat(valueText, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid Value %q: %w", item.Value, err)
		}

		currencies = append(currencies, Currency{
			NumCode:  numCode,
			CharCode: item.CharCode,
			Value:    value,
		})
	}

	sort.SliceStable(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	return currencies, nil
}
