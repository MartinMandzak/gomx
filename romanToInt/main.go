package main

import (
	"fmt"
)

var numsMap = map[int]byte{
    1: 'I',
    5: 'V',
    10: 'X',
    50: 'L',
    100: 'C',
    500: 'D',
    1000: 'M',
}

func mapLookupR2I(char byte) int{
	for key, val := range numsMap{
		if val == char{
			return key
		}
	}
	return 0
}

func romanToInt(s string) int {
	total := 0
	for idx := 0; idx< len(s)-1;idx++{
			//fmt.Printf("%c and %c \n",s[idx],s[idx+1])
			first := mapLookupR2I(s[idx])
			second := mapLookupR2I(s[idx+1])
			if first < second {
				first = first * -1
			}
			total += first
		}
	if len(s) > 0 {
      total += mapLookupR2I(s[len(s)-1])
  }
		return total
}

func main(){
	
	test:= romanToInt("D");
	fmt.Printf("%d was the result",test)
}
