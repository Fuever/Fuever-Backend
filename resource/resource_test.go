package resource

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSensitiveWords(t *testing.T) {
	for _, v := range SensitiveWords() {
		assert.True(t, v != "")
	}
}

func TestStudentMessages(t *testing.T) {
	for _, v := range StudentMessages() {
		assert.True(t, v.Number != 0)
		assert.True(t, v.Name != "")
	}
}
