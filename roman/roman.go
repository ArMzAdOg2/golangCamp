package roman

import "math"

func romanNumber(ip int) string {
	pow := 2
	result := ""
	zeroList := []string{"I", "X", "C"}
	haftList := []string{"V", "L"}
	for ip != 0 && pow > -1 {
		multi := ip / (int)(math.Pow10(pow))
		if multi == 9 {
			result += zeroList[pow] + zeroList[pow+1]
		} else if multi >= 5 {
			result += haftList[pow]
			mul := multi - 5
			for i := 0; i < mul; i++ {
				result += zeroList[pow]
			}
		} else if multi == 4 {
			result += zeroList[pow] + haftList[pow]
		} else {
			for i := 0; i < multi; i++ {
				result += zeroList[pow]
			}
		}
		ip -= multi * (int)(math.Pow10(pow))
		pow--
	}
	return result
}
