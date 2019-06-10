package phonenumber

import "strings"

func phonenumber(ip []string) map[string]int {
	result := make(map[string]int)
	for i := 0; i < len(ip); i++ {
		str := strings.Replace(ip[i], "(", "", -1)
		str = strings.Replace(str, ")", "", -1)
		str = strings.Replace(str, ",", "", -1)
		str = strings.Replace(str, " ", "", -1)
		str = strings.Replace(str, "-", "", -1)
		result[str]++
	}
	return result
}
