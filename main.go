package main

import (
	"fmt"
	"strconv"
)

func main() {
	var code uint16 = 0b0110111001010001
	fmt.Println("Incorrect Codeword:", strconv.FormatUint(uint64(code), 2))

	fmt.Println("Error should be 9:", check_codeword(code))
	code = correct_codeword(code)
	fmt.Println("Correct Codeword:  ", strconv.FormatUint(uint64(code), 2))

	fmt.Println("Error should be 0:", check_codeword(code))

}
