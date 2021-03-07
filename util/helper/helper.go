package helper

import "strconv"

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

type RequestParam string

func (p RequestParam) String() string {
	return string(p)
}

func (p RequestParam) Int64() (int64, error) {
	return strconv.ParseInt(p.String(), 10, 64)
}
