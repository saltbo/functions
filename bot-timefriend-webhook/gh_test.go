package main

import (
	"fmt"
	"testing"
)

func TestSaveToGit(t *testing.T) {
	err := SaveToGit("test")
	fmt.Println(err)
}
