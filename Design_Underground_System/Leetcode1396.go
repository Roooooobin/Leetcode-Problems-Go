package main

type checkInfo struct {
	stationName string
	time        int
}

type UndergroundSystem struct {
	// key is stationNameA + "|" + stationNameB, value is [totalTime, times], then average is totalTime / times
	stationInfo  map[string][]float64
	customerInfo map[int]checkInfo
}

func Constructor() UndergroundSystem {
	m1 := make(map[string][]float64)
	m2 := make(map[int]checkInfo)
	return UndergroundSystem{m1, m2}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {

	this.customerInfo[id] = checkInfo{stationName, t}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	checkInInfo := this.customerInfo[id]
	key := checkInInfo.stationName + "|" + stationName
	elapsed := float64(t - checkInInfo.time)
	if _, ok := this.stationInfo[key]; !ok {
		this.stationInfo[key] = []float64{elapsed, 1}
	} else {
		this.stationInfo[key] = []float64{this.stationInfo[key][0] + elapsed, this.stationInfo[key][1] + 1}
	}
	delete(this.customerInfo, id)
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	average := this.stationInfo[startStation+"|"+endStation][0] / this.stationInfo[startStation+"|"+endStation][1]
	return average
}
