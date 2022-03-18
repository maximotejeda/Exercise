package parser

// Function to parse String to identify containers
// I dont know how deep a descriptor of other people may be thats why this approach is iterative using hash
func ParseSTR(str string) bool {
	sts := map[rune]bool{
		40:  false,
		91:  false,
		123: false,
	}

	cls := map[rune]bool{
		41:  false,
		93:  false,
		125: false,
	}
	order := []rune{} // here a make a psoudo quee to last in first out
	res := false
	for _, item := range str {
		_, ok := sts[item] // Loook for the opening
		if ok {
			order = append(order, item)
		}
		_, ok = cls[item] // look for closing token
		if ok {
			if (order[len(order)-1] - item) <= 2 { // if there is a closing token
				res = true
				order = order[:len(order)-1] // pop last item in the closing search "order"
			} else {
				return false // If closing is diferent at the last opening breack
			}

		}

	}

	return res
}
