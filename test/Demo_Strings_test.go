package test

import (
	"testing"
	"strings"
)

func Test_strings_indexOf(t *testing.T) {
	var (
		fileName = "team.jpeg"
		flag = "."
	)
	index := strings.LastIndex(fileName, flag)
	if -1 == index {
		t.Error("index is -1")
	}
	t.Log(fileName[:index])
	t.Log(fileName[index + 1:])
	t.Log(index)
}
