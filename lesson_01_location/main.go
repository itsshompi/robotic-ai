package main

import "fmt"

//Modify the move function to accommodate the added
//probabilities of overshooting or undershooting
//the intended destination.

var p = []float64{0.0, 1.0, 0.0, 0.0, 0.0}
var world = []string{"green", "red", "red", "green", "green"}
var measurements = []string{"red", "green"}
var pHit = 0.6
var pMiss = 0.2
var pExact = 0.8
var pOvershoot = 0.1
var pUndershoot = 0.1

func mod(a, mod int) int {
	rem := a % mod
	if rem < 0 {
		rem += mod
	}
	return rem
}

func sense(p []float64, Z string) []float64 {
	q := []float64{}
	hit := 0.0
	for i := 0; i < len(p)-1; i++ {
		hit = 0.0
		if Z == world[i] {
			hit = 1.0
		}
		n := p[i] * (hit*pHit + (1-hit)*pMiss)
		q = append(q, n)
	}
	s := 0.0
	for _, num := range p {
		s = s + float64(num)
	}
	for i := 0; i < len(p); i++ {
		q[i] = q[i] / s
	}
	return q
}

func move(p []float64, U int) []float64 {
	q := []float64{}
	for i := 0; i < len(p); i++ {
		s := pExact * p[mod((i-U), len(p))]
		s = s + pOvershoot*p[mod((i-U-1), len(p))]
		s = s + pUndershoot*p[mod((i-U+1), len(p))]
		q = append(q, s)
	}
	return q
}

func main() {
	for i := 0; i < 1000; i++ {
		p = move(p, 1)
	}
	fmt.Printf("%v\n", p)
}
