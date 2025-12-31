package kindergarten

import (
    "errors"
    "slices"
    "strings"
)

type Cups struct {
    first string
    second string
    third string
    fourth string
}
type Garden map[string]Cups 

func NewGarden(diagram string, childrens []string) (*Garden, error) {
    lines := strings.Lines(diagram);
    var lineSlice []string
    for line := range lines {
        lineSlice = append(lineSlice, line)
    }
    
    if len(lineSlice) == 0 || lineSlice[0] != "\n" {
        return nil, errors.New("Malformed diagram")
    }
    garden := make(Garden)
    firstRow := strings.TrimRight(lineSlice[1], "\n")
    secondRow := strings.TrimRight(lineSlice[2], "\n")
    
    if len(firstRow) != len(childrens) * 2 || len(secondRow) != len(childrens) * 2 {
        return nil, errors.New("Every children, four cups")
    }
    childrenCopy := make([]string, len(childrens))
    copy(childrenCopy, childrens)
    slices.Sort(childrenCopy)

    for i := 0; i < len(childrenCopy); i++ {
        child := childrenCopy[i]
        _, exist := garden[child]
        if exist {
            return nil, errors.New("Malformed list of children (name duplicate)")
        }
        if !isValidPlant(string(firstRow[i*2])) || !isValidPlant(string(firstRow[i*2+1])) ||
           !isValidPlant(string(secondRow[i*2])) || !isValidPlant(string(secondRow[i*2+1])) {
            return nil, errors.New("Invalid Plant Code")
        }
        garden[child] = Cups{
            first:  string(firstRow[i*2]),
            second: string(firstRow[i*2+1]),
            third:  string(secondRow[i*2]),
            fourth: string(secondRow[i*2+1]),
        }
    }

	return &garden, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
    c, exist := (*g)[child]
    if exist {
    	return []string{plant(c.first), plant(c.second), plant(c.third), plant(c.fourth)}, true
    }
    return []string{}, false
}
func plant(code string) string {
    switch code {
        case "R":
        return "radishes"
        case "C":
        return "clover"
        case "G":
        return "grass"
        case "V":
        return "violets"
        default:
        panic("invalid plant")
    }
}
func isValidPlant(code string) bool {
    if code == "R" || code == "C" || code == "G" || code == "V" {
        return true
    }
    return false
}
