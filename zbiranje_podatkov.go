package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"regexp"
	"strings"
	"encoding/csv"
	"os"
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

func main() {

	csvfile, err := os.Create("rezultati.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvfile.Close()

	writer := csv.NewWriter(csvfile)

	for _, y := range years() {
		for _, d := range discipline() {

			// Katere datoteke naj izpusti
			switch d {
			case "/4x100m-relay-men": continue
			case "/4x100m-relay-women": continue
			case "/4x400m-relay-men": continue
			case "/4x400m-relay-women": continue
			}
			
			// Datoteke, ki jih želimo brati. 
			var filepath string = "rezultati" + y + "_" + d[1:] + ".html"

			// Prebere datoteko.
			data, err := ioutil.ReadFile(filepath)
			if err != nil {
				log.Fatal(err)
			}
			
			// Prebrano datoteko shrani v niz.
			var lines string = string(data)


			// Naredi poizvedbe z regularnimi izrazi.
			re := regexp.MustCompile("(?s)<tr>.+?<td class=\"col1\">.+?<span class=\"num.*?\">(?P<place>.+?)</span>.+?<td class=\"col2\">.+?<a href=\"/(?P<name>.+?)\">.+?<span class=\"picture\">.+?<span.*?>(?P<country>\\w\\w\\w)</span>.+?<td class=\"col3\">(?P<result>.+?)</td>.*?</tr>")
			matches := re.FindAllStringSubmatch(lines, -1)

			// For zanka se zapelje čez zadetke
			for _, match := range matches {
				var fileCSV []string
				disc := strings.ReplaceAll(d[1:], "-", " ")
				year := strings.ReplaceAll(y[1:], "-", " ")
				name := strings.Title(strings.ReplaceAll(match[2], "-", " "))
				country := match[3]
				r := strings.ReplaceAll(match[4], " ", "")
				result := r[2:len(r)-2]

				// Če pri mestu piše G zamenja z 1., S zamenja z 2. in B zamenja s 3.
				switch match[1]{
				case "G": fileCSV = append(fileCSV, disc, year, name, country, "1.", result)
				case "S": fileCSV = append(fileCSV, disc, year, name, country, "2.", result)
				case "B": fileCSV = append(fileCSV, disc, year, name, country, "3.", result)
				default: fileCSV = append(fileCSV, disc, year, name, country, match[1], result)
				}

				// Piše vrstice v csv datoteko
				fmt.Println("writting...")
				err := writer.Write(fileCSV)
				if err != nil {
					log.Fatal(err)
				}

				
			}
		}
	}
	writer.Flush()
}
