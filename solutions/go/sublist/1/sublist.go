package sublist

func Sublist(l1, l2 []int) Relation {
    if len(l1) == len(l2) {
        if same(l1, l2) {
            return RelationEqual
        }
        return RelationUnequal
    }
    if len(l1) == 0 {
        return RelationSublist
    }
    if len(l2) == 0 {
        return RelationSuperlist
    }

    big := l2
    little := l1
    if len(l1) > len(l2) {
        big = l1
        little = l2
    }

    sub := false
    for i:= 0; i <= len(big) - len(little); i++ {
        if same(big[i:len(little)], little) {
            sub = true
            break
        }
    }
    if sub && len(l1) > len(l2) {
        return RelationSuperlist
    }
    if sub && len(l1) < len(l2) {
        return RelationSublist
    }
    
	return RelationUnequal
}

func same(l1, l2 []int) bool {
    for i:= 0; i < len(l1); i++ {
        if l1[i] != l2[i] {
            return false
        }
    }
    return true
}
