package retina

import (
	"encoding/csv"
	"log"
	"os"
)

func LoadCsv(filename string) [][]string {
	tmpcsvfile, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error loading Retina CSV: ", err)
	}
	defer tmpcsvfile.Close()

	csvreader := csv.NewReader(tmpcsvfile)
	csvrecords, err := csvreader.ReadAll()
	if err != nil {
		log.Fatal("Error reading Retina CSV: ", err)
	}
	return csvrecords
}
