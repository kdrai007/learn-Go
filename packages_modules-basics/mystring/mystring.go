package mystring

func Reverse(str string) string {
	res := ""
	for _, s := range str {
		res = string(s) + res
	}
	return res
}
