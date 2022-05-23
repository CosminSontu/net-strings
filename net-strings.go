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

func EnsurePathPrefix(s string) string {
	if len(s) == 0 {
		return s
	}

	r0 := []rune(s)[0]

	if r0 == '/' || r0 == '?' {
		return s
	}

	return fmt.Sprintf("/%s", s)
}

func PathRewrite(source string, cutPrefix string, targetFormat string) string {
	return fmt.Sprintf(targetFormat, EnsurePathPrefix(CutPrefix(source, cutPrefix)))
}
