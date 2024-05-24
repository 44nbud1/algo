package main

// HighestPalindrome recursive func for create highest palindrome
func HighestPalindrome(s string, k int) string {

	var n = len(s)
	var chars = []rune(s)
	var changes = make([]bool, n)

	// function to make the string (s param) a palindrome
	var makePalindrome func(int, int, int) int
	makePalindrome = func(left, right, k int) int {
		if left >= right {
			return k
		}

		if chars[left] != chars[right] {
			if chars[left] > chars[right] {
				chars[right] = chars[left]
			} else {
				chars[left] = chars[right]
			}
			changes[left] = true
			changes[right] = true
			k--
		}

		if k < 0 {
			return k
		}

		return makePalindrome(left+1, right-1, k)
	}

	// function to get maximize the palindrome
	var maximizePalindrome func(int, int, int) int
	maximizePalindrome = func(left, right, k int) int {

		if left > right || k <= 0 {
			return k
		}

		if left == right {
			if chars[left] != '9' {
				chars[left] = '9'
				k--
			}
			return k
		}

		if chars[left] != '9' {
			if changes[left] && changes[right] {
				if k >= 1 {
					chars[left] = '9'
					chars[right] = '9'
					k--
				}

			} else if changes[left] || changes[right] {
				if k >= 1 {
					chars[left] = '9'
					chars[right] = '9'
					k--
				}
			} else if k >= 2 {
				chars[left] = '9'
				chars[right] = '9'
				k -= 2
			}
		}

		return maximizePalindrome(left+1, right-1, k)
	}

	// create palindrome
	k = makePalindrome(0, n-1, k)
	if k < 0 {
		return "-1"
	}

	// find max value
	k = maximizePalindrome(0, n-1, k)

	return string(chars)
}
