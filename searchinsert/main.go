
import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	for idx, el := range nums{
		if target <= el {
			return idx
		}	
	}
	return len(nums)
}

func main(){

}
