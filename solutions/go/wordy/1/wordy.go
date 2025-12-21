package wordy

import (
    "strconv"
    "strings"
)

type Operator string

const (
	OperatorPlus     Operator = "plus"
	OperatorMinus    Operator = "minus"
	OperatorDivision Operator = "divided"
	OperatorMultiple Operator = "multiplied"
    OperatorNone     Operator = "none"
)

func Answer(question string) (int, bool) {
    operands := []int{}
    operators := []Operator{}
    
    for _, term := range strings.Fields(question) {
        operand, err := strconv.Atoi(term)
        if err == nil {
            if len(operators) == len(operands) {
            	operands = append(operands, operand)
            } else {
                return 0, false
            }
            
        } else if containsOperator(term) {
            operators = append(operators, Operator(term))
        } else if strings.Contains(term, "?") {
    		term = strings.Replace(term, "?","", -1)
            operand, err := strconv.Atoi(term)
            if err == nil && len(operators) == len(operands) {    
                operands = append(operands, operand)
            } else {
                return 0, false
            }
        }
    }

    if len(operators) +1 == len(operands) {
        count := operands[0]
        for i := 0; i < len(operators); i++ {
            switch operators[i] {
                case OperatorPlus:
                count += operands[i+1]
                case OperatorMinus:
                count -= operands[i+1]
                case OperatorDivision:
                count /= operands[i+1]
                case OperatorMultiple:
                count *= operands[i+1]
            }
        }
        return count, true
    }
    return 0, false
}

func containsOperator(term string) bool {
    if term == string(OperatorMinus) || term == string(OperatorPlus) || term == string(OperatorMultiple) || term == string(OperatorDivision) {
        return true
    }
    return false
}
