package cbr

import (
	"encoding/xml"
	"fmt"
	"math"
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

type currencyDocument struct {
	XMLName xml.Name      `xml:"ValCurs"`
	Entries []currencyXML `xml:"Valute"`
}

type currencyXML struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

func Load(path string) (currencies []Currency, loadErr error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open XML file: %w", err)
	}

	defer func() {
		if err := file.Close(); err != nil && loadErr == nil {
			currencies = nil
			loadErr = fmt.Errorf("close XML file: %w", err)
		}
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var document currencyDocument
	if err := decoder.Decode(&document); err != nil {
		return nil, fmt.Errorf("decode XML: %w", err)
	}

	currencies = make([]Currency, 0, len(document.Entries))

	for _, entry := range document.Entries {
		currency, err := parseCurrency(entry)
		if err != nil {
			return nil, err
		}

		currencies = append(currencies, currency)
	}

	sort.SliceStable(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	return currencies, nil
}

func parseCurrency(entry currencyXML) (Currency, error) {
	var numCode int

	if strings.TrimSpace(entry.NumCode) != "" {
		parsedCode, err := strconv.Atoi(strings.TrimSpace(entry.NumCode))
		if err != nil {
			return Currency{}, fmt.Errorf("invalid NumCode %q: %w", entry.NumCode, err)
		}

		numCode = parsedCode
	}

	valueText := strings.ReplaceAll(strings.TrimSpace(entry.Value), ",", ".")
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return Currency{}, fmt.Errorf("invalid Value %q: %w", entry.Value, err)
	}

	if math.IsNaN(value) || math.IsInf(value, 0) {
		return Currency{}, fmt.Errorf("non-finite currency value %q", entry.Value)
	}

	return Currency{
		NumCode:  numCode,
		CharCode: entry.CharCode,
		Value:    value,
	}, nil
}
