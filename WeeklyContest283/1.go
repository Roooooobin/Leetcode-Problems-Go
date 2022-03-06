package main

func cellsInRange(s string) []string {

	res := make([]string, 0)
	rSt := s[0] - 'A'
	rEd := s[3] - 'A'
	cSt := s[1] - '0'
	cEd := s[4] - '0'
	for i := rSt; i <= rEd; i++ {
		for j := cSt; j <= cEd; j++ {
			res = append(res, string(rune(int(i)+'A'))+string(rune(int(j)+'0')))
		}
	}
	return res
}
