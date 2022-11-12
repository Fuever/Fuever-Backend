package resource

import (
	_ "embed"
	"strings"
)

//go:embed sensitive_word.dat
var sensitiveWordString string

func SensitiveWords() []string {
	return strings.Split(sensitiveWordString, "\n")
}
