package Count_All_Valid_Pickup_and_Delivery_Options

func countOrders(n int) int {

	if n == 1 {
		return 1
	}
	res, MOD := 1, int(1e9+7)
	for i := 2; i < n+1; i++ {
		res = res * (i*2 - 1) * i % MOD
	}
	return res
}
