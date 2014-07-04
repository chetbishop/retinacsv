package main

import (
	"github.com/chetbishop/retinacsv/retina"
	//"log"
	"flag"
)

var retinacsv string
var retinajob string
var outfile string

func init() {
	flag.StringVar(&retinacsv, "retinacsv", "", "Filename of Retina CSV output")
	flag.StringVar(&retinajob, "retinajobmetrics", "", "Filename of Retina CSV Job Metrics")
	flag.StringVar(&outfile, "summaryfile", "summary.csv", "Filename for summary CSV")
}

func main() {
	flag.Parse()
	csv, head := retina.LoadCsv(retinacsv)
	metrics, methead := retina.LoadJobMetrices(retinajob)
	retina.RemoveDuplicatesCsv(&csv)
	iavdetected := retina.FindIAVDetected(csv, head)
	iavcounts := retina.CountIAV(csv, head, iavdetected)
	//log.Println(metrics, methead)
	summary := retina.PercentSummary(iavcounts, metrics, methead)
	retina.WriteSummary(outfile, summary)
}
