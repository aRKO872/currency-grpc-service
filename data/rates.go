package data

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hashicorp/go-hclog"
)

type ExchangeRates struct {
	log hclog.Logger
	rates map[string]float64
}

func NewRates(l hclog.Logger) (*ExchangeRates, error) {
	er := &ExchangeRates{
		log: l,
		rates: make(map[string]float64),
	}

	err := er.getRates()

	return er, err
}

type Cubes struct {
	CubeData []Cube `xml:"Cube>Cube>Cube"`
}

type Cube struct {
	Currency string `xml:"currency,attr"`
	Rate string `xml:"rate,attr"`
}

func (e *ExchangeRates) GetRate(base, dest string) (float64, error) {
	baseRate, ok := e.rates[base]
	if !ok {
		return 0, fmt.Errorf("No rate found for currency %s", base)
	}

	destRate, ok := e.rates[dest]
	if !ok {
		return 0, fmt.Errorf("No rate found for currency %s", dest)
	}

	return destRate/baseRate, nil
}

func (e *ExchangeRates) getRates () error {
	resp, err := http.DefaultClient.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Expected 200 got %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	md := &Cubes{}
	if err := xml.NewDecoder(resp.Body).Decode(&md); err != nil {
		return fmt.Errorf("error decoding xml, %s", err.Error())
	}

	for _, cube := range md.CubeData {
		r, err := strconv.ParseFloat(cube.Rate, 64)
		if err != nil {
			return err
		}

		e.rates[cube.Currency] = r
	}

	e.rates["EUR"] = 1

	return nil
}