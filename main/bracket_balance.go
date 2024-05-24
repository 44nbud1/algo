package main

func BracketBalance(s string) string {

	// initial map to store the matching pairs of brackets
	bracketPairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// stack to store opening brackets
	var stack []rune

	for _, char := range s {
		switch char {
		case '(', '{', '[':

			// push opening brackets to stack
			stack = append(stack, char)

		case ')', '}', ']':
			// if stack is empty or top of stack is not the matching opening bracket
			if len(stack) == 0 || stack[len(stack)-1] != bracketPairs[char] {
				return "NO"
			}

			// pop the stack
			stack = stack[:len(stack)-1]
		}
	}

	// if stack is empty, all brackets are balanced
	if len(stack) == 0 {
		return "YES"
	}

	// return no
	return "NO"
}
