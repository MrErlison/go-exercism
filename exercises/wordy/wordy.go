package wordy

import (
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	question = strings.TrimRight(question, "?")
	tokens := strings.Split(question, " ")

	if tokens[0] != "What" || tokens[1] != "is" || len(tokens[2:]) == 0 {
		return 0, false
	}

	result, _ := strconv.Atoi(tokens[2])
	for i := 3; i <= len(tokens)-1; i += 2 {
		if len(tokens[i:]) < 2 {
			return result, false
		}
		operand := tokens[i]
		if tokens[i+1] == "by" {
			i++
		}
		n2, _ := strconv.Atoi(tokens[i+1])

		switch operand {
		case "plus":
			result += n2
		case "minus":
			result -= n2
		case "divided":
			result /= n2
		case "multiplied":
			result *= n2
		default:
			return 0, false
		}
	}

	return result, true
}
