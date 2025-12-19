package series

func All(n int, s string) []string {
    if n > len(s) {
        return nil
    }
    r := make([]string,0,0)
    i := 0
    for ;true; {
        r = append(r, s[i:i+n])
        i++
        if i+n > len(s) {
            break
        }
    }
	return r
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}
