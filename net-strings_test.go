package net_strings

import (
	"strings"
	"testing"
)

func TestNoSideEffects(t *testing.T) {

	path := "/a/b/c"
	prefix := "/a"
	expected := "/b/c"
	result := CutPrefix(path, prefix)

	if result != expected {
		t.Errorf("result is %s and was expecting %s", result, expected)
	} else {
		t.Logf("result %s is as expected", result)
	}
	if path == "/b/c" {
		t.Logf("path changed")
	}
	if path != "/b/c" {
		t.Logf("original string not changed")
	}
}

func TestUrlRewrite(t *testing.T) {
	sourcePath := "/a/b/c"
	cutPrefix := "/a/"
	targetPath := "/invoke/target/v1.0%s"
	expected := "/invoke/target/v1.0/b/c"
	result := PathRewrite(sourcePath, cutPrefix, targetPath)

	if expected != result {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestNoSideEffectsCut(t *testing.T) {

	path := "/a/b/c"
	prefix := "/a"
	expected := "/b/c"
	_, result, found := strings.Cut(path, prefix)

	if !found {
		t.Errorf("result is %s and was expecting %s", result, expected)
	}
}
func BenchmarkCut(b *testing.B) {
	path := "/a/b/c"
	prefix := "/a"
	for i := 0; i < b.N; i++ {
		strings.Cut(path, prefix)
	}
}
func BenchmarkSlice(b *testing.B) {
	path := "/a/b/c"
	prefix := "/a"
	for i := 0; i < b.N; i++ {
		CutPrefix(path, prefix)
	}
}

func BenchmarkEnsurePathPrefixOld(b *testing.B) {
	var inputs = []string{"abc", "", "/abc", "?abc"}

	for i := 0; i < b.N; i++ {
		ensurePathPrefixOld(inputs[i%4])
	}
}

func BenchmarkEnsurePathPrefixCurrent(b *testing.B) {
	var inputs = []string{"abc", "", "/abc", "?abc"}

	for i := 0; i < b.N; i++ {
		EnsurePathPrefix(inputs[i%4])
	}
}
