package main

import (
	"log"
	"strconv"
)

//Given an integer return an integer that is the reverse ordering of numbers
// reverseInt(15) -> 51
// reverseInt(981) -> 189
// reverseInt(500) -> 5
// reverseInt(-15) -> -51
// reverseInt(-90) -> -9

func reverseInt(num int) int {
	/*
		if num%10 == 0 {
			//get rid of zeros
			// Divide the number by 10 until it is not divisible by 10
			for num%10 == 0 && num != 0 {
				num /= 10
			}
		}
	*/
	str_num := reverse(strconv.Itoa(num))

	if num < 0 {
		// consider minus sign
		str_num = string('-') + str_num[:len(str_num)-1]
	}

	reverseNum, err := strconv.Atoi(str_num)
	if err != nil {
		log.Println("can't convert to string")
	}

	return reverseNum

}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
