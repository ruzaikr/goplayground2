package main

func isAnagram(s string, t string) bool {
	var sMap = make(map[string]int)
	for i := 0; i < len(s); i++ {
		var sChar = s[i : i+1]
		sMap[sChar]++
	}
	for i := 0; i < len(t); i++ {
		var tChar = t[i : i+1]
		sMap[tChar]--
		if sMap[tChar] == 0 {
			delete(sMap, tChar)
		}
	}
	return len(sMap) == 0
}

func main() {
	isAnagram("chicken", "kenchic")
}
