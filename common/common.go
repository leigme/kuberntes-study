package common

import (
	"os"
	"strings"
)

func HomeDir() string {
	if h := os.Getenv("HOME"); !strings.EqualFold(h, "") {
		return h
	}
	return os.Getenv("USERPROFILE")
}
