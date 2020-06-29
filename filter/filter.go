package filter

// you can use this function as GetWords filters params,
// It can change the result of GetWords,
//and there is some default filter we proper for you.
type Filter func(keyM map[string]int) map[string]int


