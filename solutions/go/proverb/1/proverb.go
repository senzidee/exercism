package proverb

import "fmt"

func Proverb(rhyme []string) []string {
    length := len(rhyme) 
	if length == 0 {
        return rhyme
    }
    var proverb []string
    for i:=0; i < length -1; i++ {
        proverb = append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
    }
    proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
    return proverb
}
