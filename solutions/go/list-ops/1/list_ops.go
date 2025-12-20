package listops

type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, i := range s {
        initial = fn(initial, i)
    }
    return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
    for i := s.Length() - 1; i >= 0; i-- {
        initial = fn(s[i], initial)
    }
    return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	result := IntList{}
    for _, i := range s {
        if fn(i) {
            result = append(result, i)
        }
    }
    return result
}

func (s IntList) Length() int {
	count := 0
    for  range s {
        count++
    }
    return count
}

func (s IntList) Map(fn func(int) int) IntList {
	result := IntList{}
    for _, i := range s {
        result = append(result, fn(i))
    }
    return result
}

func (s IntList) Reverse() IntList {
    result := IntList{}
    for i := s.Length() -1; i >= 0; i-- {
        result = append(result, s[i])
    }
	return result
}

func (s IntList) Append(lst IntList) IntList {
	result := s
    for _, item := range lst {
        result = append(result, item)
    }
    return result
}

func (s IntList) Concat(lists []IntList) IntList {
    result := s
    for _, list := range lists {
        result = result.Append(list)
    }
    return result
}
