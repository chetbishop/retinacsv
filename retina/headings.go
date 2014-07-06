package retina

func (h *IavRefHeadings) GetIavRefHeadings(slice []string) *IavRefHeadings {
	h.IAV = SlicePosition(slice, "IAV")
	h.Name = SlicePosition(slice, "Name")
	h.Patch = SlicePosition(slice, "Patch")
	return h
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
