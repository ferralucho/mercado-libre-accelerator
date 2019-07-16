package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Configuration struct {
	Name 				string `json:"name"`
}

type Site struct {
	Name 				string `json:"name"`
	CountryId 			string `json:"country_id"`
	DefaultCurrentId 	string `json:"default_currency_id"`
}

type ListingType struct {
	Id 			string `json:"id""`
	Configuration		Configuration `json:"configuration"`
	NotAvailableCategories		[]string `json:"not_available_in_categories"`
}

type Category struct {
	Id 			string `json:"id"`
	Name		string `json:"name"`
	TotalItems	int `json:"total_items_in_this_category"`
}

func main(){
	var siteID string
	var listingTypeID string

	siteID, listingTypeID, err := ingresarParametros()

	res, err := http.Get(fmt.Sprintf("https://api.mercadolibre.com/sites/%s", siteID))
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := ioutil.ReadAll(res.Body)

	var site Site

	err = json.Unmarshal(data, &site)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Site Name", site.Name)
	fmt.Println("Site Country", site.CountryId)
	fmt.Println("----------")

	res, err = http.Get(fmt.Sprintf("https://api.mercadolibre.com/sites/%s/listing_types/%s", siteID, listingTypeID))
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err = ioutil.ReadAll(res.Body)

	var listingType ListingType

	err = json.Unmarshal(data, &listingType)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Listing type Id", listingType.Id)
	fmt.Println("Listing type Name", listingType.Configuration.Name)

	fmt.Println("----------")
	mostrarCategoriasNoDisponible(listingType.NotAvailableCategories)

}

func ingresarParametros () (string, string, error) {

	fmt.Print("Ingrese el site id")
	//fmt.Scan(&siteID)
	siteID := "MLA"
	fmt.Print("Ingrese el listing type")
	//fmt.Scan(&listingTypeID)
	listingTypeID := "gold_pro"

	return siteID, listingTypeID, nil
}

func mostrarCategoriasNoDisponible(categorias []string){
	for _, categoriaNombre := range categorias {

		res, err := http.Get(fmt.Sprintf("https://api.mercadolibre.com/categories/%s", categoriaNombre))
		if err != nil {
			fmt.Println(err)
			return
		}
		data, err := ioutil.ReadAll(res.Body)

		var category Category

		err = json.Unmarshal(data, &category)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Category Id", category.Id)
		fmt.Println("Category Name", category.Name)
		fmt.Println("Category Total Items", category.TotalItems)
		fmt.Println("----------")
	}
}
