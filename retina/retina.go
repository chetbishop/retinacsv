package retina

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

type RetinaJobMetricsHeadings struct {
	JobName,
	JobFileName,
	AuditsRevision,
	ScannerVersion,
	StartTime,
	Duration,
	Credentials,
	AuditGroups,
	AddressGroups,
	IPRanges,
	HostsAttemtped,
	HostsScanned,
	NoAccess int
}

type RetinaCsvHeadings struct {
	NetBIOSName,
	DNSName,
	IP,
	MAC,
	OS,
	AuditID,
	SevCode,
	CVE,
	IAV,
	Name,
	Description,
	Date,
	Risk,
	CVSSScore,
	PCILevel,
	FixInformation,
	CCE,
	NetBIOSDomain,
	Exploit,
	Context,
	CPE,
	PCIPassFail,
	PCIReason int
}

func (h *RetinaJobMetricsHeadings) GetRetinaJobMetricsHeadings(slice []string) *RetinaJobMetricsHeadings {
	h.JobName = SlicePosition(slice, "JobName")
	h.JobFileName = SlicePosition(slice, "JobFileName")
	h.AuditsRevision = SlicePosition(slice, "AuditsRevision")
	h.ScannerVersion = SlicePosition(slice, "ScannerVersion")
	h.StartTime = SlicePosition(slice, "StartTime")
	h.Duration = SlicePosition(slice, "Duration")
	h.Credentials = SlicePosition(slice, "Credentials")
	h.AuditGroups = SlicePosition(slice, "AuditGroups")
	h.AddressGroups = SlicePosition(slice, "AddressGroups")
	h.IPRanges = SlicePosition(slice, "IPRanges")
	h.HostsAttemtped = SlicePosition(slice, "HostsAttemtped")
	h.HostsScanned = SlicePosition(slice, "HostsScanned")
	h.NoAccess = SlicePosition(slice, "NoAccess")
	return h
}

func (h *RetinaCsvHeadings) GetRetinaCsvHeadings(slice []string) *RetinaCsvHeadings {
	h.NetBIOSName = SlicePosition(slice, "NetBIOSName")
	h.DNSName = SlicePosition(slice, "DNSName")
	h.IP = SlicePosition(slice, "IP")
	h.MAC = SlicePosition(slice, "MAC")
	h.OS = SlicePosition(slice, "OS")
	h.AuditID = SlicePosition(slice, "AuditID")
	h.SevCode = SlicePosition(slice, "SevCode")
	h.CVE = SlicePosition(slice, "CVE")
	h.IAV = SlicePosition(slice, "IAV")
	h.Name = SlicePosition(slice, "Name")
	h.Description = SlicePosition(slice, "Description")
	h.Date = SlicePosition(slice, "Date")
	h.Risk = SlicePosition(slice, "Risk")
	h.CVSSScore = SlicePosition(slice, "CVSSScore")
	h.PCILevel = SlicePosition(slice, "PCILevel")
	h.FixInformation = SlicePosition(slice, "FixInformation")
	h.CCE = SlicePosition(slice, "CCE")
	h.NetBIOSDomain = SlicePosition(slice, "NetBIOSDomain")
	h.Exploit = SlicePosition(slice, "Exploit")
	h.Context = SlicePosition(slice, "Context")
	h.CPE = SlicePosition(slice, "CPE")
	h.PCIPassFail = SlicePosition(slice, "PCIPassFail")
	h.PCIReason = SlicePosition(slice, "PCIReason")
	return h
}

func SlicePosition(slice []string, value string) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}
	return -1
}

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

func LoadCsv(filename string) ([][]string, *RetinaCsvHeadings) {
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
	csvheadings := new(RetinaCsvHeadings)
	csvheadings.GetRetinaCsvHeadings(csvrecords[0])
	return csvrecords, csvheadings
}

func UniqueColumnCount(csvfile [][]string, column int) int {
	var columndata []string
	header := csvfile[0][column]
	for entry := range csvfile {
		if csvfile[entry][column] != header {
			columndata = append(columndata, csvfile[entry][column])
		}
	}
	RemoveDuplicates(&columndata)
	return len(columndata)

}
func RemoveDuplicates(xs *[]string) {
	found := make(map[string]bool)
	j := 0
	for i, x := range *xs {
		if !found[x] {
			found[x] = true
			(*xs)[j] = (*xs)[i]
			j++

		}
	}
	*xs = (*xs)[:j]
}
func RemoveDuplicatesCsv(xs *[][]string) {
	found := make(map[string]bool)
	j := 0
	for i, x := range *xs {
		y := strings.Join(x, " ")
		if !found[y] {
			found[y] = true
			(*xs)[j] = (*xs)[i]
			j++
		}

	}
	*xs = (*xs)[:j]
}
func FindIAVDetected(csvfile [][]string, headings *RetinaCsvHeadings) []string {
	var iavdetected []string
	header := csvfile[0][headings.IAV]
	for entry := range csvfile {
		if csvfile[entry][headings.IAV] != header && csvfile[entry][headings.IAV] != "N/A" {
			if strings.Contains(csvfile[entry][headings.IAV], ",") == true {
				entrysplit := strings.Split(csvfile[entry][headings.IAV], ",")
				for x := range entrysplit {
					iavdetected = append(iavdetected, entrysplit[x])
				}
			} else {
				iavdetected = append(iavdetected, csvfile[entry][headings.IAV])
			}
		}
	}
	RemoveDuplicates(&iavdetected)
	return iavdetected
}

func CountIAV(csvfile [][]string, headings *RetinaCsvHeadings, iavdetected []string) [][]string {
	var iavcounts [][]string
	for iav := range iavdetected {
		var iavcount int
		for entry := range csvfile {
			if strings.Contains(csvfile[entry][headings.IAV], iavdetected[iav]) == true {
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
