package main

func validUtf8(data []int) bool {

	cnt := 0
	for _, num := range data {
		// translate header
		if cnt == 0 {
			if (num >> 5) == 0b110 {
				cnt = 1
			} else if (num >> 4) == 0b1110 {
				cnt = 2
			} else if (num >> 3) == 0b11110 {
				cnt = 3
			} else if num>>7 != 0 {
				return false
			}
		} else {
			if (num >> 6) != 0b10 {
				return false
			}
			cnt--
		}
	}
	return cnt == 0
}
