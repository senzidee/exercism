package anagram

import (
    "sort"
    "strings"
)

func Detect(subject string, candidates []string) []string {
    var result []string
    _sLower := strings.ToLower(subject)
    _sSlice := []rune(_sLower)
    sort.Slice(_sSlice, func(i int, j int) bool { return _sSlice[i] < _sSlice[j]})
    for _, candidate := range candidates {
        _cLower := strings.ToLower(candidate)
        _cSlice := []rune(_cLower)
        sort.Slice(_cSlice, func(i int, j int) bool { return _cSlice[i] < _cSlice[j]})
        if (string(_cSlice) == string(_sSlice) && _sLower != _cLower) {
            result = append(result, candidate)
        }
    }
    
    return result
}
