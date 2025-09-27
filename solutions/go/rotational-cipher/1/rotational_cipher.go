package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
    encrypted := ""
	for _, char := range plain {
        encrypted += string(rotate(char, shiftKey))
    }
    return encrypted
}
func rotate(char rune, shiftKey int) rune {
    if char >= 65 && char <= 90 {
        char += rune(shiftKey)
        if char > 90 {
            char = char - 26
        }
    }
    if char >= 97 && char <= 122 {
        char += rune(shiftKey)
        if char > 122 {
            char = char - 26
        }
    }

    return char
}
