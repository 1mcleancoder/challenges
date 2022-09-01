package ch22_factorial

type solution1UsingRecursion struct {
}

func (s2 solution1UsingRecursion) calcFactorial(n int) int {
	return calcFactorialUsingRec(n)
}

func calcFactorialUsingRec(n int) int {
	if n == 1 {
		return 1
	}

	return n * calcFactorialUsingRec(n-1)
}

//func calcFactorialUsingRec(n, fact int) int {
//	if n == 1 {
//		return fact
//	}
//
//	return calcFactorialUsingRec(n-1, fact *= n)
//}
