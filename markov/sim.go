package main

import (
	"fmt"

	"math/rand"


	"github.com/zond/god/client"
  	"github.com/zond/god/common"
  	"github.com/zond/god/murmur"
)

func s (numTrains int, safetyTime int, rideTime int, minTransTime int, avgTransTime int) []int{
	d := make([]int,0)
	
	trains := make([]int,numTrains)
	c := 0
	avg := float64(avgTransTime - minTransTime)
	t := 0
	dep := minTransTime + int(rand.ExpFloat64() * avg)
	t = dep
	trains[c] = dep + rideTime
	c = (c + 1) % numTrains
	d = append(d, dep)
	for t < 43200 { //in seconds
		if trains[c] > t {
			t = trains[c]
		}
		t += minTransTime + int(rand.ExpFloat64() * avg)
		if t < dep + safetyTime{
			t = dep + safetyTime
		}
		dep = t
		d = append(d,dep)
	}
	fmt.Println(d)
	return d
}

func main(){
	a := (s(3,30,50,20, 27))
	fmt.Println(a)
	fmt.Println(a)
	conn := client.MustConn("localhost:9191")
	bytes := common.MustJSONEncode(a)
	conn.Put(murmur.HashString("hi"), bytes)
}
		
	
