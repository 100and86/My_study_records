package main

//二分查找
import (
	"fmt"
)

func main() {
	sli := []int{1, 3, 7, 9, 12, 14}
	goal1:=13
	goal2:=12
	fmt.Println(binarysearch(sli,goal1))
	fmt.Println(binarysearch(sli,goal2))


}
func binarysearch(sli []int, goal int) int {

	left := 0
	right := len(sli) - 1
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		if sli[mid] == goal {
			fmt.Println("yes")
			return mid
		} else if sli[mid] < goal {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Println("No")
	return -1
}
