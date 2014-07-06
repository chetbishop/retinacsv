package retina

import (
	"encoding/csv"
	"log"
	"os"
)

func LoadJobMetrices(filename string) ([][]string, *RetinaJobMetricsHeadings) {
	tmpcsvfile, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error loading Retina Job Metrics: ", err)
	}
	defer tmpcsvfile.Close()

	csvreader := csv.NewReader(tmpcsvfile)
	csvrecords, err := csvreader.ReadAll()
	if err != nil {
		log.Fatal("Error reading Retina Job Metrics: ", err)
	}
	csvheadings := new(RetinaJobMetricsHeadings)
	csvheadings.GetRetinaJobMetricsHeadings(csvrecords[0])
	return csvrecords, csvheadings
}

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
