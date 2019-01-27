package handler

import (
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	s, _ := os.Getwd()
	print(s)

}
