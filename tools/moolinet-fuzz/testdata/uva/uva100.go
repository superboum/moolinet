package main

import "fmt"

var cache map[int]int

func algo(n int) int {
	if _, ok := cache[n]; !ok {
		if n == 1 {
			cache[n] = 1
		} else if n%2 == 1 {
			cache[n] = 1 + algo(3*n+1)
		} else {
			cache[n] = 1 + algo(n/2)
		}
	}
	return cache[n]
}

func main() {
	cache = make(map[int]int)
	var n, m, a, b int

	for {
		nb, err := fmt.Scanf("%d %d", &n, &m)
		if nb == 0 || err != nil {
			break
		}

		a, b = n, m
		if a > b {
			a, b = b, a
		}

		res := 0
		for i := a; i <= b; i++ {
			cur := algo(i)
			if cur > res {
				res = cur
			}
		}

		fmt.Printf("%d %d %d\n", n, m, res)
	}
}
