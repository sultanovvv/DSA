package leetcode

func longestCommonPrefix(strs []string) string {

	cnum := 0

	for i := 0; ; i++ {
		cnum++
		comm := ""
		for _, str := range strs {
			if len(str) < i+1 {
				cnum--
				return string([]rune(str)[:cnum])
			}
			if comm == "" {
				comm = string([]rune(str)[i])
			} else if comm != string([]rune(str)[i]) {
				cnum--
				return string([]rune(str)[:cnum])
			}
		}
	}
}
