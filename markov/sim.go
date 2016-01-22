package main

import (
	"fmt"

	"math/rand"

	"container/list"
)

func s (int numTrains, int safetyTime, int rideTime, int minTransTime, int avgTransTime){
	var trains[numTrains] int
	c := 0
	avg : = avgTransTime - minTransTime
	t := 0
	dep := minTransTime + int(ExpFloat64() * avg)
	t = dep
	trains[c] = dep + rideTime
	c = (c + 1) % numTrains
	fmt.println(dep)
	for t < 43200 { //in seconds
		if trains[c] > t {
			t = trains[c]
		}
		t += minTransTime + int(ExpFloat64() * avg)
		if t < dep + safetyTime{
			t = dep + safteyTime
		}
		dep = t
		fmt.println(dep)
	}
}
		
	
