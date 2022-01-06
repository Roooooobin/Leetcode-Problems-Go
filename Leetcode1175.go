package main

func numPrimeArrangements(n int) int {

	nPrime := 0
	MOD := 1000000007
	for i := 2; i <= n; i++ {
		if judgePrime(i) {
			nPrime++
		}
	}
	return calcFactorial(nPrime, MOD) * calcFactorial(n-nPrime, MOD) % MOD
}

func judgePrime(x int) bool {

	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func calcFactorial(x int, mod int) int {

	res := 1
	for i := 2; i <= x; i++ {
		res = res * i % mod
	}
	return res
}
