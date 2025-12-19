package sublist

func Sublist(l1, l2 []int) Relation {
    switch {
        case len(l1) == len(l2) && contains(l1, l2):
        	return RelationEqual
        case len(l1) > len(l2) && contains(l1, l2):
        	return RelationSuperlist
        case len(l1) < len(l2) && contains(l2, l1):
        	return RelationSublist
        default:
        	return RelationUnequal
    }
}

func contains(big, little []int) bool {
    if len(little) == 0 {
        return true
    }
    for i:= 0; i <= len(big) - len(little); i++ {
        if isSame(big[i:len(little)], little) {
            return true
        }
    }
    return false
}

func isSame(l1, l2 []int) bool {
    for i:= range l1 {
        if l1[i] != l2[i] {
            return false
        }
    }
    return true
}
