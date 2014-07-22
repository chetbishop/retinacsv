package retina

import (
	"encoding/csv"
	"log"
	"os"
)

//LoadCsv will read a CSV file into a [][]string.
func LoadCsv(filename string) [][]string {
	tmpcsvfile, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error loading CSV: ", err)
	}
	defer tmpcsvfile.Close()

	csvreader := csv.NewReader(tmpcsvfile)
	csvrecords, err := csvreader.ReadAll()
	if err != nil {
		log.Fatal("Error reading CSV: ", err)
	}
	return csvrecords
}
