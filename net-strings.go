package net_strings

import (
	"fmt"
	"strings"
)

type net_string string

func CutPrefix(s string, prefix string) string {

	if strings.Index(s, prefix) == 0 {
		s = string(s[len(prefix):])
	}
	return s
}

func EnsurePrefix(s string, prefix string) string {
	if (strings.Index(s, prefix)) == 0 {
		return s
	} else {
		return fmt.Sprintf("%s%s", prefix, s)
	}
}

func PathRewrite(source string, cutPrefix string, targetFormat string) string {
	return fmt.Sprintf(targetFormat, EnsurePrefix(CutPrefix(source, cutPrefix), "/"))
}
