package main

import(
	"fmt"
	"encoding/csv"
	"os"
	"log"
)

type Tekmovalec struct {
	ime string
	drzava string
	rezultat string
	disciplina string
	olimpijske string
}

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

func readCSV(filename string) ([][]string, error) {

	// Odpri CSV datoteko
	dat, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer dat.Close()

	// Preberi datoteko v spremenljivko
	vrstice, err := csv.NewReader(dat).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return vrstice, nil
}

func main() {

	lines, err := readCSV("rezultati.csv")
	if err != nil {
		log.Fatal(err)
	}

	// For zanka po vrsticah v CSV datoteki
	for _, line := range lines {
		fmt.Println(line)
	}

}




