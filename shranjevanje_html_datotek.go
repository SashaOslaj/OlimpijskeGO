package main

import (
	"fmt"
	"os"
	"net/http"
	"io"
)

func years() []string {
	// Vrne seznam vseh poletnih olimpijskih iger.
	var year []string = []string{
		"/rio-2016", "/london-2012", "/beijing-2008", "/athens-2004",
		"/sydney-2000", "/atlanta-1996", "/barcelona-1992", "/seoul-1988",
		"/los-angeles-1984", "/moscow-1980", "/montreal-1976", "/munich-1972",
		"/mexico-1968", "/tokyo-1964", "/rome-1960", "/melbourne-stockholm-1956",
		"/helsinki-1952", "/london-1948", "/berlin-1936", "/los-angeles-1932",
		"/amsterdam-1928", "/paris-1924", "/antwerp-1920", "/stockholm-1912",
		"/london-1908", "/st-louis-1904", "/paris-1900", "/athens-1896",
	}

	return year
}

func sport() string {
	// Vrne niz atletika, ker gledamo le to disciplino.
	sport := "/athletics"

	return sport
}

func discipline() []string {
	// Vrne seznam vseh disciplin atletike
	var disciplines []string = []string{
		"/10000m-men", "/100m-men", "/110m-hurdles-men", "/1500m-men",
		"/200m-men", "/20km-walk-men", "/3000m-steeplechase-men",
		"/400m-hurdles-men", "/400m-men",

		"/4x100m-relay-men", "/4x400m-relay-men",
		"/5000m-men", "/50km-walk-men", "/800m-men",
		"/decathlon-men", "/discus-throw-men", "/hammer-throw-men",
		"/high-jump-men", "/javelin-throw-men", "/long-jump-men",
		"/marathon-men", "/pole-vault-men", "/shot-put-men", "/triple-jump-men",

		"/10000m-women", "/100m-hurdles-women", "/100m-women",
		"/1500m-women", "/200m-women", "/20km-race-walk-women",
		"/3000m-steeplechase-women", "/400m-hurdles-women", "/400m-women",
		"/4x100m-relay-women", "/4x400m-relay-women",
		"/5000m-women",
		"/800m-women", "/discus-throw-women", "/hammer-throw-women",
		"/heptathlon-women", "/high-jump-women", "/javelin-throw-women",
		"/long-jump-women", "/marathon-women", "/pole-vault-women",
		"/shot-put-women", "/triple-jump-women",
	}

	return disciplines
}

func saveHTML(filepath string, url string) (err error) {
	
	// Ustvari datoteko
	out, err := os.Create(filepath)
	if err != nil {
		// Preveri, če datoteka že obstaja
		fmt.Println(err)
		return err
	}
	defer out.Close()

	// Pridobi podatke iz url-ja.
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	// Zapiši datoteko.
	_, err = io.Copy(out, resp.Body)
	if err != nil  {
		fmt.Println(err)
	  	return err
	}
  
	return nil

}

func main() {

	// Osnovni url naslov
	var mainURL string = "https://www.olympic.org"

	for _, y := range years() {
		for _, d := range discipline() {
			
			// Ustvari url naslov in ime datoteke 
			var url string = mainURL + y + sport() + d
			var filepath string = "rezultati" + y + "_" + d[1:] + ".html"
			
			// Preveri, če datoteka že obstaja, če ne jo ustvari
			if _, err := os.Stat(filepath); err != nil {
				fmt.Printf("Saving...%q \n", url)
				saveHTML(filepath, url)
			} else {
				fmt.Printf("Already exists...%q \n", url)
			}			
			
		}
	}

}
