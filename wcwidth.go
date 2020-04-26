package wcwidth


type IntSet struct {
	set map[int]bool
}

func (set *IntSet) Add(v int) {
	set.set[v] = true
}

func (set *IntSet) Delete(v int) {
	delete(set.set, v)
}

func (set *IntSet) Exist(v int) bool {
	_, found := set.set[v]
	return found
}

var zeroWidthCF *IntSet

func NewIntSet(arr []int) *IntSet {
	set := &IntSet{
		set: make(map[int]bool),
	}
	for _, v := range arr {
		set.Add(v)
	}
	return set
}


func bisearch(ucs int, table [][2]int, ubound int) int {
	lbound := 0

	if ucs < table[0][0] || ucs > table[ubound][1] {
		return 0
	}

	for ubound >= lbound {
		mid := (lbound + ubound) / 2
		if ucs > table[mid][1] {
			lbound = mid + 1
		} else if ucs < table[mid][0] {
			ubound = mid - 1
		} else {
			return 1
		}
	}

	return 0
}

// width val is the unicode point of a character
// for example: ✨ is 0x2728
func width(ucs int) int {
	if zeroWidthCF.Exist(ucs) {
		return 0
	}

	// C0/C1 control character
	if ucs < 32 || (0x07F <= ucs && ucs < 0x0A0) {
		return -1
	}

	if bisearch(ucs, zeroWidth, len(zeroWidth)-1) == 1 {
		return 0
	}

	return 1 + bisearch(ucs, widthEastasion, len(widthEastasion)-1)
}

func StringWidth(input string) int {
	runes := []rune(input)
	ret := 0
	for _, ucs := range runes {
		wcw := width(int(ucs))
		if wcw < 0 {
			return -1
		}
		ret += wcw
	}
	return ret
}


func init() {
	zeroWidthCF = NewIntSet([]int{
		0,       // Null (Cc)
		0x034F,  // Combining grapheme joiner (Mn)
		0x200B,  // Zero width space
		0x200C,  // Zero width non-joiner
		0x200D,  // Zero width joiner
		0x200E,  // Left-to-right mark
		0x200F,  // Right-to-left mark
		0x2028,  // Line separator (Zl)
		0x2029,  // Paragraph separator (Zp)
		0x202A,  // Left-to-right embedding
		0x202B,  // Right-to-left embedding
		0x202C,  // Pop directional formatting
		0x202D,  // Left-to-right override
		0x202E,  // Right-to-left override
		0x2060,  // Word joiner
		0x2061,  // Function application
		0x2062,  // Invisible times
		0x2063,  // Invisible separator
	})
}
