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
	return sli[i].Year < sli[j].Year
}

func (sli sortTrack) Swap(i, j int) {
	sli[i], sli[j] = sli[j], sli[i]
}
func (sli sortTrack) String() string {
	buf := new(bytes.Buffer)
	for _, s := range sli {
		fmt.Fprintf(buf, "%s\n", fmt.Sprint(*s))
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
	sort.Sort(tracks) //直接修改切片的顺序
	fmt.Print(tracks)
	//书中的例子比较古老了，现在实际开发中，经常直接用，Less函数根据需要自定义
	// `倒序`
	// 	sort.Slice(tracks, func(i, j int) bool {
	// 	return tracks[i].Artist > tracks[j].Artist
	// })
	// `多级排序`
	// sort.Slice(tracks, func(i, j int) bool {
	// 	x, y := tracks[i], tracks[j]

	// 	if x.Title != y.Title {
	// 		return x.Title < y.Title
	// 	}
	// 	if x.Year != y.Year {
	// 		return x.Year < y.Year
	// 	}
	// 	return x.Length < y.Length
	// })
	sort.Slice(tracks, func(i, j int) bool { //传入一个切片，go 1.8之后加入的
		return tracks[i].Year < tracks[j].Year
	})
	//这就不需要自己定义：
	// type sortTrack []*Track
	// func Len()
	// func Less()
	// func Swap()
	// Go 1.21 之后:slices.SortFunc,按比较函数升序排序，且它不是稳定排序；稳定版本可以用 slices.SortStableFunc
	// slices.SortFunc(tracks, func(a, b *Track) int {
	// 	return cmp.Compare(a.Year, b.Year)
	// })
}

func TestPalindrome(t *testing.T) {
	sli := Palindrome{1, 2, 3, 2, 1}
	if IsPalindrome(sli) {
		fmt.Print("yes")
		return
	}
	fmt.Print("No")

}

type Palindrome []int

func (p Palindrome) Len() int {
	return len(p)
}
func (p Palindrome) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p Palindrome) Swap(i, j int) {
	panic("禁止使用")
}

func IsPalindrome(s sort.Interface) bool {
	n := s.Len()
	j := n - 1
	for i := 0; i < n/2; i++ {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
		j--

	}
	return true
}
