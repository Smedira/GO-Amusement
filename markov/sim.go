package main

import (
	"fmt"

	"math/rand"
)

func s (numTrains int, safetyTime int, rideTime int, minTransTime int, avgTransTime int) int{
	cc:=1
	trains := make([]int,numTrains)
	c := 0
	avg := float64(avgTransTime - minTransTime)
	t := 0
	dep := minTransTime + int(rand.ExpFloat64() * avg)
	t = dep
	trains[c] = dep + rideTime
	c = (c + 1) % numTrains
	fmt.Println(dep)
	for t < 43200 { //in seconds
		if trains[c] > t {
			t = trains[c]
		}
		t += minTransTime + int(rand.ExpFloat64() * avg)
		if t < dep + safetyTime{
			t = dep + safetyTime
		}
		dep = t
		fmt.Println(dep)
		cc += 1
	}
	return cc
}

func main(){
	a := s(3,30,50,20, 27)
	fmt.Println(a)
}
		
	
