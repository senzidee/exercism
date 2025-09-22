package strand

func ToRNA(dna string) string {
    rna := ""
	runes := []rune(dna)
    length := len(runes)
    for i := 0; i < length; i++ {
        switch runes[i] {
            case 'A':
            	rna += "U"
            case 'C':
            	rna += "G"
            case 'G':
            	rna += "C"
            case 'T':
            	rna += "A"
        }
    }

    return rna
}
