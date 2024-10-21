
package main

import (
	"fmt"
)
/*
func contains(num int, slice []int) bool{
	for _, value := range slice{
		if value == num{
			return true
		}
	}
	return false
}

func removeElement(val int, slice []int) []int {
	for idx, element := range slice{
		if val == element {
			return append(slice[:idx], slice[idx+1:]...)
		}
	}
	return slice
}

func singleNumber(nums []int) int {
	var result = make([]int,0)

	for _,val := range nums{
		if contains(val, result){
			result = removeElement(val, result)
		} else{
			result = append(result, val)
		}
	}

	return result[0]
}
*/

func singleNumber(nums []int) int {
	var res = 0
	for _, num := range nums{
		//fmt.Println(res)    // omg that's so clever, it's adding the numbers on the binary and ^= XORing them
		res^=num
	}
	return res
}

func main(){
	nums := []int{4,1,2,1,2}
	var res = singleNumber(nums)
	fmt.Println(res)	

}
