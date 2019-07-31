package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v1"
	"sync"
)

const (
	SITES_ALL_URL             string = "https://api.mercadolibre.com/sites"
	SITES_SEARCH_URL          string = "https://api.mercadolibre.com/sites/%s"
	CURRENCIES_CONVERSION_URL string = "https://api.mercadolibre.com/currency_conversions/search?from=%s&to=%s"
	CURRENCY_USD              string = "USD"
)

type Sites []struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Site struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	CountryId         string `json:"country_id"`
	DefaultCurrencyId string `json:"default_currency_id"`
}

type CurrencyConversion struct {
	From  string
	To    string
	Ratio float64
}

func main() {
	currencies, err := GetAllCurrencies()
	if err != nil {
		fmt.Println(fmt.Sprintf("Error when trying to get currencies: %v", err))
		return
	}
	result, _ := json.Marshal(currencies)
	fmt.Println(string(result))
}

func GetAllCurrencies() (map[string]float64, error) {
	sites, err := GetAllSites()

	if err != nil {
		return nil, err
	}

	c := make(chan *CurrencyConversion)
	defer close(c)
	result := make(map[string]float64)

	var wg sync.WaitGroup

	go HandleResults(c, &wg, len(*sites), result)

	for _, site := range *sites {
		wg.Add(1)
		go ProcessSiteId(c, site.Id)
	}
	wg.Wait()
	return result, nil
}

func HandleResults(c chan *CurrencyConversion, wg *sync.WaitGroup, loop int, m map[string]float64) {
	if wg == nil || loop == 0 {
		return
	}

	var conv *CurrencyConversion

	for i := 0; i < loop; i++ {
		conv = <-c
		if conv != nil {
			m[conv.From] = conv.Ratio
		}
		// Notificamos que se ha procesado otro elemento del WaitGroup.
		wg.Done()
	}
}

func ProcessSiteId(c chan *CurrencyConversion, siteId string) {
	convRate, err := GetCurrencyConversion(siteId, CURRENCY_USD)
	if err != nil {
		c <- nil
		return
	}
	c <- convRate
}

// Busca el ratio de conversión entre dos monedas haciendo uso de la API de currencies de MercadoLibre.
// API de ejemplo: https://api.mercadolibre.com/currency_conversions/search?from=ARS&to=USD
func GetCurrencyConversion(siteId, to string) (*CurrencyConversion, error) {
	site, err := GetSite(siteId)
	if err != nil {
		return nil, err
	}

	fmt.Println(fmt.Sprintf("About to convert from '%s' to '%s'.", site.DefaultCurrencyId, to))
	uri := fmt.Sprintf(CURRENCIES_CONVERSION_URL, site.DefaultCurrencyId, to)
	resp, err := resty.R().Get(uri)
	if err != nil {
		return nil, err
	}

	var conversion CurrencyConversion
	err = json.Unmarshal(resp.Body(), &conversion)
	if err != nil {
		return nil, err
	}
	fmt.Println(fmt.Sprintf("Convertion rate from '%s' to '%s' successfully obtained.", site.DefaultCurrencyId, to))
	conversion.From = site.DefaultCurrencyId
	conversion.To = to

	return &conversion, nil
}

//https://api.mercadolibre.com/sites
func GetAllSites() (*Sites, error) {
	resp, err := resty.R().Get(SITES_ALL_URL)
	if err != nil {
		return nil, err
	}
	var sites Sites
	err = json.Unmarshal(resp.Body(), &sites)
	if err != nil {
		return nil, err
	}
	fmt.Println(fmt.Sprintf("Found sites: %d", len(sites)))
	return &sites, nil
}

func GetSite(siteId string) (*Site, error) {
	fmt.Println(fmt.Sprintf("About to get siteId '%s'...", siteId))
	resp, err := resty.R().Get(fmt.Sprintf(SITES_SEARCH_URL, siteId))
	if err != nil {
		return nil, err
	}
	var site Site
	err = json.Unmarshal(resp.Body(), &site)
	if err != nil {
		return nil, err
	}
	fmt.Println(fmt.Sprintf("SiteId '%s' successfully obtained.", siteId))
	return &site, nil
}
