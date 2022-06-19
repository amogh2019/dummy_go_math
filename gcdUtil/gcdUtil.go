package gcdutil

func gcd(a, b int) int {

	if b > a {
		return gcd(b, a)
	}
	if b == 0 {
		return a
	}
	if b == 1 {
		return 1
	}
	return gcd(a%b, b)
}
