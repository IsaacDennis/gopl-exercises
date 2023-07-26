package exercises

import "strings"

func VariadicJoin(sep string, strs ...string) string {
	return strings.Join(strs, sep)
}
