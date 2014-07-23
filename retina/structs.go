package retina

//ScanCsv is a struct to hold information from the Retina scan.
type ScanCsv struct {
	ScanData           [][]string
	JobMetrics         [][]string
	IavRef             [][]string
	ScanDataHeadings   RetinaCsvHeadings
	JobMetricsHeadings RetinaJobMetricsHeadings
	IavRefHeadings     IavRefHeadings
}

//CsvAnalysis is a struct to hold analysis information about a Retina scan.
type CsvAnalysis struct {
	IavDetected       []string
	IavDetails        []IavDetails
	IavDetailsSummary [][]string
	DevicesDetected   []string
	DeviceDetails     []Device
	Summary           [][]string
}

//Device is a struct that holds information about a device found by a Retina scan.
type Device struct {
	DeviceIP   string
	DeviceName string
	Iav        []string
	IavCount   int
}

//IavCounts is a struct with details for each specific IAV.
type IavDetails struct {
	Iav      string
	DeviceIP []string
	Count    int
}

//IavRefHeadings is a struct that contains the column number for headings of the IAV reference file.
type IavRefHeadings struct {
	IAV,
	Name,
	Patch int
}

//RetinaJobMetricsHeadings is a struct that contains the column number for headings of the Retina Job Metrics file.
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

//RetinaCsvHeadings is a struct that contains the column number for the headings of the Retina CSV.
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
