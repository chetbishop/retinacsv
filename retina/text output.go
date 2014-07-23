package retina

import (
	//"encoding/csv"
	//"log"
	"os"
	//"strconv"
	//"strings"
)

//IavDetailsOut creates a text file for each IAV and lists the IP addresses found to be vulnerable.
func (analysis *CsvAnalysis) IavDetailsOut() {
	err := os.Mkdir("IAV", 0660)
	if err != nil {
		return
	}
	for _, iav := range analysis.IavDetails {
		iavfile, err := os.OpenFile("IAV/"+iav.Iav+".txt", os.O_RDWR|os.O_CREATE, 0660)
		if err != nil {
			return
		}
		defer iavfile.Close()
		iavfile.WriteString("IAV Found: " + iav.Iav + "\r\n")
		for _, ip := range iav.DeviceIP {
			iavfile.WriteString(ip + "\r\n")
		}
	}
}

//DeviceDetails creates a text file for each device and lists the IAVs found.
func (analysis *CsvAnalysis) DeviceDetailsOut() {
	err := os.Mkdir("Device", 0660)
	if err != nil {
		return
	}
	for _, devicefound := range analysis.DeviceDetails {
		var filename string
		if devicefound.DeviceName != "" {
			filename = devicefound.DeviceName
		} else {
			filename = devicefound.DeviceIP
		}
		devicefile, err := os.OpenFile("Device/"+filename+".txt", os.O_RDWR|os.O_CREATE, 0660)
		if err != nil {
			return
		}
		defer devicefile.Close()
		devicefile.WriteString("Device Name: " + devicefound.DeviceName + "\r\n")
		devicefile.WriteString("Device IP: " + devicefound.DeviceIP + "\r\n")
		for _, iav := range devicefound.Iav {
			devicefile.WriteString(iav + "\r\n")
		}
	}
}
