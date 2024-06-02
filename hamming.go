package main

//  TODO: 15, 11 hamming code generalize later

func check_codeword(code uint16) uint16 {
	var bits uint16 = code
	var error uint16 = 0
	var i uint16
	bits >>= 1
	for i = 1; i < 16; i++ {
		if bits&1 != 0 {
			error ^= i
		}
		bits >>= 1
	}
	return error
}

func correct_codeword(code uint16) uint16 {
	error := check_codeword(code)
	for start := uint16(1); start < 16; start <<= 1 {
		if error&start > 0 {
			code ^= (1 << start)
		}
	}
	return code
}
