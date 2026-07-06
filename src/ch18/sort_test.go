package sort_test

import (
	"bytes"
	"fmt"
	"sort"
	"testing"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
}
type sortTrack []*Track

func (sli sortTrack) Len() int {
	return len(sli)
}
func (sli sortTrack) Less(i, j int) bool {
	return sli[i].Title < sli[j].Title
}

func (sli sortTrack) Swap(i, j int) {
	sli[i], sli[j] = sli[j], sli[i]
}
func (sli sortTrack) String() string {
	buf := new(bytes.Buffer)
	for _, s := range sli {
		fmt.Fprintf(buf,"%s\n",fmt.Sprint(*s))
	}
	return buf.String()
}

func TestSort(t *testing.T) {
	var tracks = sortTrack{
		{"Go", "Delilah", "From the Roots Up", 2012},
		{"Go", "Moby", "Moby", 1992},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011},
	}
	sort.Sort(tracks)
	fmt.Print(tracks)
}
