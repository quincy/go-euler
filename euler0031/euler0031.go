/*
In England the currency is made up of pound, £, and pence, p, and there are
eight coins in general circulation:

1p, 2p, 5p, 10p, 20p, 50p, £1 (100p) and £2 (200p).
It is possible to make £2 in the following way:

1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p
How many different ways can £2 be made using any number of coins?
*/

package main

import (
	"fmt"
)

const (
	onePence    = 1
	twoPence    = 2
	fivePence   = 5
	tenPence    = 10
	twentyPence = 20
	fiftyPence  = 50
	onePound    = 100
	twoPounds   = 200
)

func sumMoney(a, b, c, d, e, f, g, h int) int {
	return a*onePence + b*twoPence + c*fivePence + d*tenPence + e*twentyPence + f*fiftyPence + g*onePound + h*twoPounds
}

func main() {
	m := make(map[string]bool)
	var sum int

	for a := 0; a <= 200; a++ {
		sum = a
		if sum == twoPounds {
			k := fmt.Sprintf("2£ = %d·1p", a)
			m[k] = true
			break
		}
		if sum+onePence > twoPounds {
			break
		}
		for b := 0; b <= 100; b++ {
			sum = a*onePence + b*twoPence
			if sum == twoPounds {
				k := fmt.Sprintf("2£ = %d·1p %d·2p", a, b)
				m[k] = true
				break
			}
			if sum+twoPence > twoPounds {
				break
			}
			for c := 0; c <= 40; c++ {
				sum = a*onePence + b*twoPence + c*fivePence
				if sum == twoPounds {
					k := fmt.Sprintf("2£ = %d·1p %d·2p %d·5p", a, b, c)
					m[k] = true
					break
				}
				if sum+fivePence > twoPounds {
					break
				}
				for d := 0; d <= 20; d++ {
					sum = a*onePence + b*twoPence + c*fivePence + d*tenPence
					if sum == twoPounds {
						k := fmt.Sprintf("2£ = %d·1p %d·2p %d·5p %d·10p", a, b, c, d)
						m[k] = true
						break
					}
					if sum+tenPence > twoPounds {
						break
					}
					for e := 0; e <= 10; e++ {
						sum = a*onePence + b*twoPence + c*fivePence + d*tenPence + e*twentyPence
						if sum == twoPounds {
							k := fmt.Sprintf("2£ = %d·1p %d·2p %d·5p %d·10p %d·20p", a, b, c, d, e)
							m[k] = true
							break
						}
						if sum+twentyPence > twoPounds {
							break
						}
						for f := 0; f <= 4; f++ {
							sum = a*onePence + b*twoPence + c*fivePence + d*tenPence + e*twentyPence + f*fiftyPence
							if sum == twoPounds {
								k := fmt.Sprintf("2£ = %d·1p %d·2p %d·5p %d·10p %d·20p %d·50p", a, b, c, d, e, f)
								m[k] = true
								break
							}
							if sum+fiftyPence > twoPounds {
								break
							}
							for g := 0; g <= 2; g++ {
								sum = a*onePence + b*twoPence + c*fivePence + d*tenPence + e*twentyPence + f*fiftyPence + g*onePound
								if sum == twoPounds {
									k := fmt.Sprintf("2£ = %d·1p %d·2p %d·5p %d·10p %d·20p %d·50p %d·1£", a, b, c, d, e, f, g)
									m[k] = true
									break
								}
								if sum+onePound > twoPounds {
									break
								}
								for h := 0; h <= 1; h++ {
									sum = a*onePence + b*twoPence + c*fivePence + d*tenPence + e*twentyPence + f*fiftyPence + g*onePound + h*twoPounds
									if sum == twoPounds {
										k := fmt.Sprintf("2£ = %d·1p %d·2p %d·5p %d·10p %d·20p %d·50p %d·1£ %d·2£", a, b, c, d, e, f, g, h)
										m[k] = true
									}
								}
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(len(m))
}
