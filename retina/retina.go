package retina

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

//LoadScan will create a ScanCsv struct from a Retina scan.
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
	var iavdetected []string
	header := scanStruct.ScanData[0][scanStruct.ScanDataHeadings.IAV]
	for entry := range scanStruct.ScanData {
		iav := scanStruct.ScanData[entry][scanStruct.ScanDataHeadings.IAV]
		if iav != header && iav != "N/A" {
			var q string
			if strings.Contains(iav, ",") == true {
				entrysplit := strings.Split(iav, ",")
				for x := range entrysplit {
					q = entrysplit[x]
				}
			} else {
				q = iav
			}
			iavdetected = append(iavdetected, q)
		}
	}
	analysis.IavDetected = iavdetected
	RemoveDuplicates(&analysis.IavDetected)
}

func (analysis *CsvAnalysis) GetIavDetails(scanStruct *ScanCsv) {
	var iavcountsout []IavDetails
	iavdetected := analysis.IavDetected
	for _, iavfound := range iavdetected {
		var iavcount int
		var count IavDetails
		for entry := range scanStruct.ScanData {
			iaventry := scanStruct.ScanData[entry][scanStruct.ScanDataHeadings.IAV]
			if strings.Contains(iaventry, iavfound) == true {
				count.DeviceIP = append(count.DeviceIP, scanStruct.ScanData[entry][scanStruct.ScanDataHeadings.IP])
				iavcount++
			}
		}

		count.Iav = iavfound
		count.Count = iavcount
		iavcountsout = append(iavcountsout, count)
	}
	analysis.IavDetails = iavcountsout
}
func (analysis *CsvAnalysis) GetDeviceList(scanStruct *ScanCsv) {
	var devicelist []string
	header := scanStruct.ScanData[0][scanStruct.ScanDataHeadings.IP]
	for entry := range scanStruct.ScanData {
		device := scanStruct.ScanData[entry][scanStruct.ScanDataHeadings.IP]
		if device != header && device != "N/A" {
			devicelist = append(devicelist, device)
		}
	}
	analysis.DevicesDetected = devicelist
	RemoveDuplicates(&analysis.DevicesDetected)
}

func (analysis *CsvAnalysis) GetDeviceDetails(scanStruct *ScanCsv) {
	var devicedetails []Device
	devicelist := analysis.DevicesDetected
	for _, device := range devicelist {
		var devicecount int
		var devicestruct Device
		for entry := range scanStruct.ScanData {
			deviceentry := scanStruct.ScanData[entry][scanStruct.ScanDataHeadings.IP]
			if strings.Contains(deviceentry, device) == true {
				if scanStruct.ScanData[entry][scanStruct.ScanDataHeadings.IAV] != "N/A" {
					devicestruct.Iav = append(devicestruct.Iav, scanStruct.ScanData[entry][scanStruct.ScanDataHeadings.IAV])
					devicestruct.DeviceName = scanStruct.ScanData[entry][scanStruct.ScanDataHeadings.NetBIOSName]
					devicecount++
				}
			}
		}

		devicestruct.DeviceIP = device
		devicestruct.IavCount = devicecount
		devicedetails = append(devicedetails, devicestruct)
	}
	analysis.DeviceDetails = devicedetails
}

func (analysis *CsvAnalysis) PercentSummary(scanStruct *ScanCsv) {
	var summary [][]string
	summary = append(summary, []string{"IAV", "Number of Hosts Found Vulnerable", "Number of Hosts Found", "Number of Hosts Compliant", "Percentage Compliant"})
	for _, iav := range analysis.IavDetails {
		var writestring []string
		numiavfound := iav.Count
		numhosttotal, _ := strconv.Atoi(scanStruct.JobMetrics[1][scanStruct.JobMetricsHeadings.HostsScanned])
		numhostscompliant := numhosttotal - numiavfound
		percent := strconv.FormatFloat(float64(numhostscompliant)/float64(numhosttotal), 'f', 5, 32)
		writestring = []string{iav.Iav, strconv.Itoa(numiavfound), strconv.Itoa(numhosttotal), strconv.Itoa(numhostscompliant), percent}
		summary = append(summary, writestring)
	}
	analysis.Summary = summary
}

func (analysis *CsvAnalysis) WriteSummary(fileoutname string) {
	summaryfile, err := os.OpenFile(fileoutname, os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		log.Fatal("Error creating summary: ", err)
	}
	defer summaryfile.Close()
	summarywriter := csv.NewWriter(summaryfile)
	err = summarywriter.WriteAll(analysis.Summary)
	if err != nil {
		log.Fatal("Error writing to summary CSV: ", err)
	}
}
