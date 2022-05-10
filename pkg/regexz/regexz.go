package regexz

import (
	"regexp"
)

// RegexpExtract 正则抽取
func RegexpExtract(regex, src, alias string) string {
	var result []byte
	pattern := regexp.MustCompile(regex)
	for _, sub := range pattern.FindAllStringSubmatchIndex(src, -1) {
		result = pattern.ExpandString(result, alias, src, sub)
	}

	return string(result)
}
