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
	order := []rune{} // here i make a psoudo queue to last in first out
	res := false
	for _, item := range str {
		_, ok := sts[item] // Loook for the opening
		if ok {
			order = append(order, item)
		}
		_, ok = cls[item] // look for closing token
		if ok {
			if len(order) == 0 { // if theres a closing tag and no opening in the list
				return false
			}
			if (item-order[len(order)-1]) <= 2 && (item-order[len(order)-1]) > 0 { // if there is a closing token
				res = true
				order = order[:len(order)-1] // pop last item in the closing search "order"
			} else {
				return false // If closing is diferent at the last opening breack
			}

		}

	}
	if len(order) > 0 {
		return false
	}

	return res
}

func Rparser(val string, token []rune, ret bool) bool {
	if token == nil {
		token = []rune{}
	}
	open := map[rune]bool{
		'(': false,
		'{': false,
		'[': false,
	}

	close := map[rune]bool{
		')': false,
		'}': false,
		']': false,
	}
	if token == nil {
		token = []rune{}
	}
	for i, item := range val {
		_, ok := open[item]
		if ok {
			token = append(token, item)
			//			fmt.Println(token)
			return Rparser(val[i+1:], token, true) && ret
		}
		_, ok = close[item]
		if ok {
			if len(token) == 0 {
				return false
			}
			token = token[:len(token)-1]
			return Rparser(val[i+1:], token, true) && ret
		} else {
			continue
		}

	}
	if len(token) > 0 {
		return false
	}
	return ret
}
