package q5

import "strings"


func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = removerAlfanumerico(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}

	}
	return true
}

func eAlfanumerico(rs rune) bool {

	if (rs >= 'a' && rs <= 'z') || (rs >= '0' && rs <= '9') {
		return true
	}
	return false

}

func removerAlfanumerico(s string) string {
	var resultado strings.Builder

	for _, rs := range s {
		if eAlfanumerico(rs) {
			resultado.WriteRune(rs)
		}
	}
	return resultado.String()
}
