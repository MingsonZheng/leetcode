package main

import "fmt"

// 剑指 Offer 16. 数值的整数次方

func main() {
	fmt.Println(myPow(2.00000, 10))
	fmt.Println(myPow(2.10000, 3))
	fmt.Println(myPow(2.00000, -2))
}

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return rPow(x, n)
	} else {
		return 1 / (rPow(x, -1*(n+1)) * x)
	}
}

func rPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	halfPow := rPow(x, n/2)
	if n%2 == 1 {
		return halfPow * halfPow * x
	} else {
		return halfPow * halfPow
	}
}
