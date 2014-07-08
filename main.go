package main

import (
	"flag"
	"github.com/chetbishop/retinacsv/retina"
	"log"
)

var scanname string
var outfile string
var iavref string

func init() {
	flag.StringVar(&scanname, "scanname", "servers", "Filename of Retina CSV output")
	flag.StringVar(&outfile, "summary", "summary.csv", "Filename for summary CSV")
	flag.StringVar(&iavref, "iavref", "iavdetails.csv", "File containing information related it IAVs")
}

func main() {
	flag.Parse()
	ScanCsv := retina.LoadScan(scanname, iavref)
	retina.RemoveDuplicatesCsv(&ScanCsv.ScanData)
	anaysisStruct := new(retina.CsvAnalysis)
	anaysisStruct.FindIAVDetected(ScanCsv)
	anaysisStruct.CountIAV(ScanCsv)
	log.Println(anaysisStruct)
	//summary := retina.PercentSummary(iavcounts, metrics, methead)
	//retina.WriteSummary(outfile, summary)
}
