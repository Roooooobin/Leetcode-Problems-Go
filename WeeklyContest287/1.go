package main

func convertTime(current string, correct string) int {

	convert := func(s string) int {
		return 600*int(s[0]-'0') + 60*int(s[1]-'0') + 10*int(s[3]-'0') + int(s[4]-'0')
	}
	diff := convert(correct) - convert(current)
	res := 0
	res += diff / 60
	diff = diff % 60
	res += diff / 15
	diff = diff % 15
	res += diff / 5
	diff = diff % 5
	res += diff
	return res
}
