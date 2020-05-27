type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
    return len(s)
}

func SortChars(s string) string {
    r := []rune(s)
    sort.Sort(sortRunes(r))
    return string(r)
}

func contains(anagrams [][]string, key string) bool {
    for i := 0; i < len(anagrams); i++ {
        for k := 0; k < len(anagrams[i]); k++ {
            if anagrams[i][k] == key {
                return true
            }
        }
    }
    return false
}

func groupAnagrams(strs []string) [][]string {
    cpy := make([]string, len(strs))
    anagrams := make([][]string, 0)
    for i := 0; i < len(strs); i++ {
        cpy[i] = strs[i]
        cpy[i] = SortChars(cpy[i])
    }
    for i := 0; i < len(cpy); i++ {
        if contains(anagrams, strs[i]) {
            continue
        }
        anagrams = append(anagrams, []string{ strs[i] })
        for k := i+1; k < len(cpy); k++ {
            if cpy[i] == cpy[k] {
                anagrams[len(anagrams)-1] = append(anagrams[len(anagrams)-1], strs[k])
            }
        }
    }
    return anagrams
}
