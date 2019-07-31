func partition(s string) [][]string {
	res := make([][]string, 0)
	tmp := make([]string, 0)
	dfs(s, 0, &res, &tmp)
	return res
}

func dfs(s string, index int, res *[][]string, tmp *[]string) {
    if index == len(s) {
        temp := make([]string, 0)
        temp = append(temp, (*tmp)...)
        *res = append(*res, temp)
        return
    }

    for i := index; i < len(s); i++ {
        *tmp = append(*tmp, s[index:i+1])
        dfs(s, i+1, res, tmp)
    }
}

func isPalindrome(s string) bool {
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

// 以上代码可以找全排列