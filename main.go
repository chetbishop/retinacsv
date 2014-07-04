package main

import (
	"flag"
	"github.com/chetbishop/retinacsv/retina"
)

var retinacsv string
var retinajob string
var outfile string

func init() {
	flag.StringVar(&retinacsv, "csv", "", "Filename of Retina CSV output")
	flag.StringVar(&retinajob, "job", "", "Filename of Retina CSV Job Metrics")
	flag.StringVar(&outfile, "summary", "summary.csv", "Filename for summary CSV")
}

func main() {
	flag.Parse()
	csv, head := retina.LoadCsv(retinacsv)
	metrics, methead := retina.LoadJobMetrices(retinajob)
	retina.RemoveDuplicatesCsv(&csv)
	iavdetected := retina.FindIAVDetected(csv, head)
	iavcounts := retina.CountIAV(csv, head, iavdetected)
	summary := retina.PercentSummary(iavcounts, metrics, methead)
	retina.WriteSummary(outfile, summary)
}
