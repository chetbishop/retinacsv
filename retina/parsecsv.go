package retina

import (
	"strings"
)

//UniqueColumnCount is not currently in use
//func UniqueColumnCount(csvfile [][]string, column int) int {
//	var columndata []string
//	header := csvfile[0][column]
//	for entry := range csvfile {
//		if csvfile[entry][column] != header {
//			columndata = append(columndata, csvfile[entry][column])
//		}
//	}
//	RemoveDuplicates(&columndata)
//	return len(columndata)
//}
//RemoveDuplicates will remove the duplicates in a CsvAnalysis struct that has a type of []string.
func (a *CsvAnalysis) RemoveDuplicates() {
	found := make(map[string]bool)
	xs := a.IavDetected
	j := 0
	for i, x := range xs {
		if !found[x] {
			found[x] = true
			(xs)[j] = (xs)[i]
			j++

		}
	}
	xs = (xs)[:j]
	a.IavDetected = xs
}

//RemoveDuplicates removes the duplicate entries in a []string.
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

//RemoveDuplicatesCsv removes duplicate lines from a CSV file.
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
