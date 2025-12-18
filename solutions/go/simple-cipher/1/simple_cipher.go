package cipher

import (
    "strings"
    "unicode"
)

type shift string
type vigenere string

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if (distance == 0 || distance < -25 || distance > 25) {
        return nil
    }
    return shift(rune(97+(distance%26 + 26)%26))
}

func (c shift) Encode(input string) string {
	return shiftString(input, string(c), true)
}

func (c shift) Decode(input string) string {
	return shiftString(input, string(c), false)
}

func NewVigenere(key string) Cipher {
    if len(key) < 2 || !isValidKey(key) {
        return nil
    }
	return vigenere(key)
}

func (v vigenere) Encode(input string) string {
	return shiftString(input, string(v), true)
}

func (v vigenere) Decode(input string) string {
	return shiftString(input, string(v), false)
}

func shiftString(input string, key string, encode bool) string {
    var result strings.Builder
    keyIndex := 0
    keyLength := len(key)
	for _, r := range input {
		if unicode.IsLetter(r) {
            distance := int(key[keyIndex] - 97)
            if !encode {
                distance = -distance
            }
			r = rune((int(unicode.ToLower(r)- 97)+distance+26)%26 + 97)
			result.WriteRune(r)
            keyIndex = (keyIndex + 1) % keyLength
		}
	}

	return result.String()
}
func isValidKey(key string) bool {
	unique := make(map[rune]bool)
	for _, r := range key {
		if !unicode.IsLower(r) {
			return false
		}
		unique[r] = true
	}

	return len(unique) > 1
}