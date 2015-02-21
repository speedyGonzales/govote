package govote

import (
//"sort"
)

//get elemenet's position in array of strings
func getIndexOf(a string, list []string) int {
	for p, el := range list {
		if el == a {
			return p
		}
	}
	return -1
}

/*func compare(s1, s1 string) int {
	if s1 == s1 {
		return 0
	}
	if s1 > s2 {
		return 1
	} else {
		return -1
	}
}*/
