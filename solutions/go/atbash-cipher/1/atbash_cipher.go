package atbash

func Atbash(s string) string {
    var result []rune
    i := 1
    for _, c := range s {
        if i == 6 {
            result = append(result, ' ')
            i = 1
        }
        if c >= 97 && c <= 122 {
            result = append(result, rune(219 - int(c)))
            i++
        } else if c >= 65 && c <= 90 {
            result = append(result, rune(122 - (int(c) -65)))
            i++
        } else if  c >= 48 && c <= 57 {
            result = append(result, c)
            i++
        }
    }
	if result[len(result) -1] == ' ' {
        return string(result[:len(result) -1])
    }
    return string(result)
}
