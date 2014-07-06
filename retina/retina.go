package retina

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

func LoadScan(filename string, iavRefFile string) *ScanCsv {
	scanStruct := new(ScanCsv)
	scanStruct.ScanData = LoadCsv(filename + ".csv")
	scanStruct.ScanDataHeadings.GetRetinaCsvHeadings(scanStruct.ScanData[0])
	scanStruct.JobMetrics = LoadCsv(filename + "_JobMetrics.csv")
	scanStruct.JobMetricsHeadings.GetRetinaJobMetricsHeadings(scanStruct.JobMetrics[0])
	scanStruct.IavRef = LoadCsv(iavRefFile)
	scanStruct.IavRefHeadings.GetIavRefHeadings(scanStruct.IavRef[0])
	return scanStruct
}

func (analysis *CsvAnalysis) FindIAVDetected(scanStruct *ScanCsv) {
	var iavdetected []IavDetect
	header := scanStruct.ScanData[0][scanStruct.ScanDataHeadings.IAV]
	for entry := range scanStruct.ScanData {
		iav := scanStruct.ScanData[entry][scanStruct.ScanDataHeadings.IAV]
		if iav != header && iav != "N/A" {
			var q IavDetect
			if strings.Contains(iav, ",") == true {
				entrysplit := strings.Split(iav, ",")
				for x := range entrysplit {
					q.Iav = entrysplit[x]
				}
			} else {
				q.Iav = iav
			}
			iavdetected = append(iavdetected, q)
		}
	}
	analysis.IavDetected = iavdetected
	analysis.RemoveDuplicates()

}

func (scanStruct *ScanCsv) CountIAV(iavdetected []string) [][]string {
	var iavcounts [][]string
	for iav := range iavdetected {
		var iavcount int
		for entry := range scanStruct.ScanData {
			iaventry := scanStruct.ScanData[entry][scanStruct.ScanDataHeadings.IAV]
			if strings.Contains(iaventry, iavdetected[iav]) == true {
				iavcount++
			}
		}
		iavcounts = append(iavcounts, []string{iavdetected[iav], strconv.Itoa(iavcount)})
	}
	return iavcounts
}
func PercentSummary(iavcounts [][]string, jobmetricsfile [][]string, jobmetricshead *RetinaJobMetricsHeadings) [][]string {
	var summary [][]string
	summary = append(summary, []string{"IAV", "Number of Hosts Found Vulnerable", "Number of Hosts Found", "Number of Hosts Compliant", "Percentage Compliant"})
	for x := range iavcounts {
		var writestring []string
		numiavfound, _ := strconv.Atoi(iavcounts[x][1])
		numhosttotal, _ := strconv.Atoi(jobmetricsfile[1][jobmetricshead.HostsScanned])
		numhostscompliant := numhosttotal - numiavfound
		percent := strconv.FormatFloat(float64(numhostscompliant)/float64(numhosttotal), 'f', 2, 32)
		writestring = []string{iavcounts[x][0], iavcounts[x][1], jobmetricsfile[1][jobmetricshead.HostsScanned], strconv.Itoa(numhostscompliant), percent}
		summary = append(summary, writestring)
	}
	return summary
}

func WriteSummary(fileoutname string, csvfile [][]string) {
	summaryfile, err := os.OpenFile(fileoutname, os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal("Error creating summary: ", err)
	}
	defer summaryfile.Close()
	summarywriter := csv.NewWriter(summaryfile)
	err = summarywriter.WriteAll(csvfile)
	if err != nil {
		log.Fatal("Error writing to summary CSV: ", err)
	}
}
