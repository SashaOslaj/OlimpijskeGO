package main

import(
	"fmt"
	"encoding/csv"
	"os"
	"log"
	"sort"
	"strconv"
	"strings"
)

// Tekmovalec je ...
type Tekmovalec struct {
	ime string
	drzava string
	mesto string
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

	// Odpre in prebere CSV datoteko
	lines, err1 := readCSV("rezultati.csv")
	if err1 != nil {
		log.Fatal(err1)
	}

	// Odprla bom nekaj datotek in v njih zapisovala v for zanki

	// Datoteka za drzave
	drzaveFile, err2 := os.Create("poizvedbe/drzave.txt")
	if err2 != nil {
		log.Fatal(err2)
	}

	// Datoteka za stevilo olimpijskih iger pri vseh disciplinah
	oiFile, err3 := os.Create("poizvedbe/oi.txt")
	if err3 != nil {
		log.Fatal(err3)
	}

	rezultati := make(map[string][]*Tekmovalec)

	// For zanka po vrsticah v CSV datoteki
	for _, line := range lines {

		// Shrani Tekmovalca v spremenljivko tekm
		var tekm *Tekmovalec = &Tekmovalec{line[2],line[3],line[4],line[5],line[0],line[1]}

		// Preveri, ce je nek tekmovalec ze v slovarju, kjer so kluci imena tekmovalcev
		if _, ok := rezultati[line[2]]; ok {
			// Pri nekem kljucu doda tekmo v seznam
			rezultati[line[2]] = append(rezultati[line[2]],tekm)
		} else {
			// V slovarju ustvari kljuc z imenom tekmovalca in prvo tekmo, ki jo najde v csv datoteki
			rezultati[line[2]] = []*Tekmovalec{tekm}
		}
	}

	// Slovarji s pomocjo katerih bom naredila nekaj poizvedb
	slovarDisciplinPrvaMesta := make(map[string][][]string)
	slovarDrzav := make(map[string]int)

	
	for _, value := range rezultati {

		// Naredi slovar, kjer so kljuci krtaice drzav in vrednosti so stevila koliko razlicnih tekmovalcev prihaja iz te drzave.
		if _,ok := slovarDrzav[value[0].drzava]; ok {
			slovarDrzav[value[0].drzava] ++
		} else {
			slovarDrzav[value[0].drzava] = 1
		}


		for _, r := range value {
		
			oi := r.olimpijske[len(r.olimpijske)-4:]
			mesto := r.mesto
			rez := r.rezultat
			sez := []string{oi, mesto, rez}


			// Preveri, ce je disciplina ze v slovarju, sicer naredi nov kljuc discipline.
			if _, ok := slovarDisciplinPrvaMesta[r.disciplina]; ok {
				if mesto == "1." || mesto == "2." || mesto == "3." {
					// Pri nekem kljucu doda tekmo v seznam, ce je bil tekmovalec prvi
					slovarDisciplinPrvaMesta[r.disciplina] = append(slovarDisciplinPrvaMesta[r.disciplina], []string{oi, mesto, rez})
				} 
			} else {
				if mesto == "1." || mesto == "2." || mesto == "3." {
					// V slovarju ustvari kljuc z disciplino in prvo tekmo, ki jo najde v slovarju rezultati
					slovarDisciplinPrvaMesta[r.disciplina] = [][]string{sez}
				} 
			}
		}
	}

	for key, value := range slovarDrzav {
		v := int64(value)
		d,_ := drzaveFile.WriteString("Iz drzave s kratico " + key + " je bilo " + strconv.FormatInt(v, 10) + " tekmovalcev. \n" )
		fmt.Println(d)
	}

	for key, value := range slovarDisciplinPrvaMesta {

		v := int64(len(value))
		dis,_ := oiFile.WriteString("Discipline " + key + " was on " + strconv.FormatInt(v, 10) + " olympic games. \n")
		fmt.Println(dis)
		//fmt.Println(key, value)
		d := key

		k := strings.Replace(key, " ", "_", -1)

		disciplineFile, err4 := os.Create(fmt.Sprintf("poizvedbe/%s.txt", k))
		if err4 != nil {
			log.Fatal(err4)
		}

		prvatrimesta := make(map[string][][]string)

		for _, i := range value {

			l := i[0]
			m := i[1]
			r := i[2]
			s := []string{m, r}
			
			if _,ok := prvatrimesta[l]; ok {
				prvatrimesta[l] = append(prvatrimesta[l], []string{m, r})
			} else {
				prvatrimesta[l] = [][]string{s}
			}
		}
		for key, value := range prvatrimesta{
			sort.Slice(value, func(i, j int) bool {	return value[i][0] < value[j][0] })

			fmt.Println(key, value)

			if len(value) > 3 {
				if len(value[0]) != 2 || len(value[1]) != 2 || len(value[2]) != 2 || len(value[3]) != 2 {
					continue
				} else {
					disc, _ := disciplineFile.WriteString(d + ", " + key + ", " + value[0][0] + " " + value[0][1] + ", " + value[1][0] + " " + value[1][1] + ", " + value[2][0] + " " + value[2][1] + ", " + value[3][0] + " " + value[3][1] + "\n")
					fmt.Println(disc, value)
				}
			} else if len(value) > 2 {
				if len(value[0]) != 2 || len(value[1]) != 2 || len(value[2]) != 2 {
					continue
				} else {
					disc, _ := disciplineFile.WriteString(d + ", " + key + ", " + value[0][0] + " " + value[0][1] + ", " + value[1][0] + " " + value[1][1] + ", " + value[2][0] + " " + value[2][1] + "\n")
					fmt.Println(disc, value)
				}
			} else if len(value) > 1 {
				if len(value[0]) != 2 || len(value[1]) != 2 {
					continue
				} else {
					disc, _ := disciplineFile.WriteString(d + ", " + key + ", " + value[0][0] + " " + value[0][1] + ", " + value[1][0] + " " + value[1][1] + "\n")
					fmt.Println(disc, value)
				}
			} else {
				if len(value[0]) != 2 {
					continue
				} else {
					disc, _ := disciplineFile.WriteString(d + ", " + key + " " + value[0][0] + " " + value[0][1] + "\n")
					fmt.Println(disc, value)
				}
			}
		}
		disciplineFile.Close()
	}
	drzaveFile.Close()
	oiFile.Close()
}