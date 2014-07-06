package retina

type ScanCsv struct {
	ScanData           [][]string
	JobMetrics         [][]string
	IavRef             [][]string
	ScanDataHeadings   RetinaCsvHeadings
	JobMetricsHeadings RetinaJobMetricsHeadings
	IavRefHeadings     IavRefHeadings
}
type CsvAnalysis struct {
	IavDetected []IavDetect
	IavCounts   []IavCounts
	Summary     [][]string
}
type IavDetect struct {
	Iav string
}
type IavCounts struct {
	Iav   string
	Count int
}
type IavRefHeadings struct {
	IAV,
	Name,
	Patch int
}
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
