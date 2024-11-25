package reverse

func Reverse(input string) string {
	res := []rune{}
	runes := []rune(input)
	for i:= len(runes) - 1; i >= 0; i--{
		res = append(res,runes[i])
	}
	return string(res)
}
