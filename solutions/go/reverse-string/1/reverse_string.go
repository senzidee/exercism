package reverse

func Reverse(input string) string {
	var reversed string
    runes := []rune(input)
    for i := len(runes) -1; i >= 0; i-- {
        reversed += string(runes[i])
    }
    return reversed
}
