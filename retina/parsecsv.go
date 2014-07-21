package retina

import (
	"strings"
)

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
