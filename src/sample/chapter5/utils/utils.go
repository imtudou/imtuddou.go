package utils

import "time"

func Cal(n1 float64, n2 float64, operator byte) float64 {

	var res float64
	switch operator {
	case '+':
		res = n1 + n2
		break
	case '-':
		res = n1 - n2
		break
	case '*':
		res = n1 * n2
		break
	case '/':
		res = n1 / n2
		break
	default:
	}
	return res
}

func Now() time.Time {
	return time.Now()
}
