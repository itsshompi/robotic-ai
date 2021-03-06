package main

import "fmt"

var p = []float64{0.2, 0.2, 0.2, 0.2, 0.2}
var world = []string{"green", "red", "red", "green", "green"}
var measurements = []string{"red", "green"}
var motions = []int{1, 1}
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
	for i := 0; i < len(p); i++ {
		hit := 0.0
		if Z == world[i] {
			hit = 1.0
		}
		n := p[i] * (hit*pHit + (1-hit)*pMiss)
		q = append(q, n)
	}
	s := 0.0
	for _, num := range q {
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
	for k := 0; k < len(measurements); k++ {
		p = sense(p, measurements[k])
		p = move(p, motions[k])
	}
	fmt.Printf("%v\n", p)
}
