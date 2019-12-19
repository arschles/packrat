package file

// WordSet is a set data structure for Words
//
// It returns two functions. The first is to add a word to the set, the second is to
// get all the words in the set
func WordSet() (func(...Word), func() Words) {
	set := map[Word]struct{}{}
	setter := func(words ...Word) {
		for _, c := range words {
			set[c] = struct{}{}
		}
	}
	getter := func() Words {
		ret := []Word{}
		for c := range set {
			ret = append(ret, c)
		}
		return ret
	}
	return setter, getter
}
