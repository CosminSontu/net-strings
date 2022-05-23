package net_strings

import (
	"fmt"
	"strings"
)

func CutPrefix(s string, prefix string) string {

	if strings.Index(s, prefix) == 0 {
		s = string(s[len(prefix):])
	}
	return s
}

func ensurePathPrefixOld(s string) string {
	if strings.TrimSpace(s) == "" {
		return ""
	}

	r0 := []rune(s)[0]

	if r0 == '/' || r0 == '?' {
		return s
	}

	return fmt.Sprintf("/%s", s)
}

func EnsurePathPrefix(s string) string {
	if strings.TrimSpace(s) == "" {
		return ""
	}

	r0 := []rune(s[:1])[0]

	if r0 == '/' || r0 == '?' {
		return s
	}

	sb := strings.Builder{}
	sb.Grow(len(s) + 1)
	sb.WriteString("/")
	sb.WriteString(s)

	return sb.String()
}

func PathRewrite(source string, cutPrefix string, targetFormat string) string {
	return fmt.Sprintf(targetFormat, EnsurePathPrefix(CutPrefix(source, cutPrefix)))
}
