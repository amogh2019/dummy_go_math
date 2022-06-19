package gcdutil

func FindGcd(a, b int) int {

	if b > a {
		return FindGcd(b, a)
	}
	if b == 0 {
		return a
	}
	if b == 1 {
		return 1
	}
	return FindGcd(a%b, b)
}
